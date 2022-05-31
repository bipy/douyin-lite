package controllers

import (
	"douyin-lite/app/models"
	"douyin-lite/app/queries"
	"douyin-lite/pkg/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"sync"
)

func FavoriteAction(c echo.Context) error {
	curID, err := Authorize(c)
	if err != nil {
		return err
	}
	params := c.QueryParams()
	videoID, err := strconv.Atoi(params.Get("video_id"))
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse("Illegal video_id"))
	}
	actionType, err := strconv.Atoi(params.Get("action_type"))
	if err != nil || actionType != 1 && actionType != 2 {
		return c.JSON(http.StatusOK, utils.FailResponse("Illegal action_type"))
	}
	if actionType == 1 {
		_ = queries.DouyinDB.DoFavorite(curID, videoID)
	} else if actionType == 2 {
		_ = queries.DouyinDB.CancelFavorite(curID, videoID)
	}
	return c.JSON(http.StatusOK, utils.SuccessResponse(echo.Map{}))
}

func GetFavoriteList(c echo.Context) error {
	curID, err := Authorize(c)
	if err != nil {
		return err
	}

	userID, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse("Illegal user_id"))
	}

	// videos
	videos, err := queries.DouyinDB.GetFavoriteList(userID)
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse(err.Error()))
	}

	videoIDs := make([]int, len(videos))
	for i, v := range videos {
		videoIDs[i] = v.Id
	}

	videoMap := map[int]*models.Video{}
	for i := range videos {
		videoMap[videos[i].Id] = &videos[i]
	}

	// authors
	userIDs := make([]int, len(videos))
	for i, v := range videos {
		userIDs[i] = v.AuthorId
	}

	var users []models.User
	var curFollow []int
	var uErr, foErr error

	// sync
	wg := sync.WaitGroup{}
	wg.Add(2)

	// user
	go func() {
		defer wg.Done()
		users, uErr = queries.DouyinDB.GetUserInfos(userIDs)
	}()

	// follow
	go func() {
		defer wg.Done()
		curFollow, foErr = queries.DouyinDB.GetFollows(curID, userIDs)
	}()

	wg.Wait()

	if uErr != nil || foErr != nil {
		return c.JSON(http.StatusOK, utils.FailResponse("Get Data Failed"))
	}

	// user
	userMap := map[int]*models.User{}
	for i := range users {
		userMap[users[i].Id] = &users[i]
	}

	// follow
	for _, f := range curFollow {
		userMap[f].IsFollow = true
	}

	// link
	for i := range videos {
		videos[i].Author = userMap[videos[i].AuthorId]
		videos[i].IsFavorite = true
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse(echo.Map{
		"video_list": videos,
	}))
}
