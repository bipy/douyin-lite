package controllers

import (
	"douyin-lite/app/models"
	"douyin-lite/app/queries"
	"douyin-lite/pkg/configs"
	"douyin-lite/pkg/repository"
	"douyin-lite/pkg/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func GetFeed(c echo.Context) error {
	// params
	curID, _ := Authorize(c)

	latestTime, _ := strconv.ParseInt(c.QueryParam("latest_time"), 10, 64)
	nextTime := latestTime

	if latestTime == 0 {
		nextTime = time.Now().UnixMilli()
		latestTime = nextTime
	}

	latestTime -= repository.MaxFeedBackwardsMilliSec

	// videos
	videos, err := queries.DouyinDB.GetFeed(latestTime / 1000)
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse(err.Error()))
	}

	if len(videos) == 0 {
		return c.JSON(http.StatusOK, utils.SuccessResponse(echo.Map{
			"next_time": latestTime,
		}))
	}

	nextTime = videos[0].CreateTime * 1000

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
	var curFollow, curFavorite []int
	var uErr, foErr, faErr error
	if curID != -1 {
		// sync
		wg := sync.WaitGroup{}
		wg.Add(3)

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

		// favorite
		go func() {
			defer wg.Done()
			curFavorite, faErr = queries.DouyinDB.GetFavorite(curID, videoIDs)
		}()

		wg.Wait()
	} else {
		users, uErr = queries.DouyinDB.GetUserInfos(userIDs)
	}

	if uErr != nil || foErr != nil || faErr != nil {
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

	// favorite
	for _, f := range curFavorite {
		videoMap[f].IsFavorite = true
	}

	// link
	for i := range videos {
		videos[i].Author = userMap[videos[i].AuthorId]
	}

	for i := range videos {
		videos[i].CoverUrl = fmt.Sprintf("%sfile/%s", configs.URLPrefix, videos[i].CoverUrl)
		videos[i].PlayUrl = fmt.Sprintf("%sfile/%s", configs.URLPrefix, videos[i].PlayUrl)
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse(echo.Map{
		"next_time":  nextTime,
		"video_list": videos,
	}))
}
