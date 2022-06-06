package controllers

import (
	"douyin-lite/app/models"
	"douyin-lite/app/queries"
	"douyin-lite/pkg/repository"
	"douyin-lite/pkg/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"sync"
)

func PublishAction(c echo.Context) error {
	curID, err := MustAuthorize(c)
	if err != nil {
		return err
	}
	title := c.FormValue("title")
	if title == "" || len(title) > repository.MaxVideoTitleLength {
		return c.JSON(http.StatusOK, utils.FailResponse("Empty Title"))
	}
	// TODO Storage
	coverUrl := "http://192.168.2.102/sample-cover"
	playUrl := "http://192.168.2.102/sample-video"
	_, err = queries.DouyinDB.CreateVideo(curID, playUrl, coverUrl, title)
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, utils.SuccessResponse(echo.Map{}))
}

func GetPublishList(c echo.Context) error {
	curID, err := MustAuthorize(c)
	if err != nil {
		return err
	}

	userID, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse("Illegal user_id"))
	}

	// videos
	videos, err := queries.DouyinDB.GetPublishList(userID)
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse(err.Error()))
	}

	if len(videos) > 0 {
		videoIDs := make([]int, len(videos))
		for i, v := range videos {
			videoIDs[i] = v.Id
		}

		videoMap := map[int]*models.Video{}
		for i := range videos {
			videoMap[videos[i].Id] = &videos[i]
		}

		var user *models.User
		isFollow := false
		var curFavorite []int
		var uErr, foErr, faErr error

		// sync
		wg := sync.WaitGroup{}
		wg.Add(3)

		// user
		go func() {
			defer wg.Done()
			user, uErr = queries.DouyinDB.GetUserInfo(userID)
		}()

		// follow
		go func() {
			defer wg.Done()
			isFollow, foErr = queries.DouyinDB.IsFollow(curID, userID)
		}()

		// favorite
		go func() {
			defer wg.Done()
			curFavorite, faErr = queries.DouyinDB.GetFavorite(curID, videoIDs)
		}()

		wg.Wait()

		if uErr != nil || foErr != nil || faErr != nil {
			return c.JSON(http.StatusOK, utils.FailResponse("Get Data Failed"))
		}

		// follow
		user.IsFollow = isFollow

		// favorite
		for _, f := range curFavorite {
			videoMap[f].IsFavorite = true
		}

		// link
		for i := range videos {
			videos[i].Author = user
		}
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse(echo.Map{
		"video_list": videos,
	}))
}
