package controllers

import (
	"douyin-lite/app/models"
	"douyin-lite/app/queries"
	"douyin-lite/pkg/repository"
	"douyin-lite/pkg/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"io/ioutil"
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
	data, err := c.FormFile("data")
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse(err.Error()))
	}

	f, err := data.Open()
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse(err.Error()))
	}

	id := uuid.New().String()

	all, err := ioutil.ReadAll(f)
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse(err.Error()))
	}

	err = ioutil.WriteFile("file/play/"+id, all, 0644)
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse(err.Error()))
	}

	err = utils.MakeSnapshot(id)
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse(err.Error()))
	}

	_, err = queries.DouyinDB.CreateVideo(curID, "play/"+id, "cover/"+id, title)
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
