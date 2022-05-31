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
	"time"
)

func GetHello(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello World!")
}

func GetFeed(c echo.Context) error {
	// params
	curID := 1
	latestTime, _ := strconv.ParseInt(c.QueryParam("latest_time"), 10, 32)
	nextTime := latestTime
	if latestTime == 0 {
		nextTime = time.Now().Unix()
		latestTime = nextTime - repository.MaxFeedBackwards
	}
	// videos
	videos, err := queries.DouyinDB.GetFeed(latestTime)
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse(err.Error()))
	}
	if len(videos) != 0 {
		nextTime = int64(videos[0].CreateTime)
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
	// sync
	wg := sync.WaitGroup{}
	wg.Add(3)
	var users []models.User
	var curFollow, curFavorite []int
	var uErr, foErr, faErr error
	// user
	go func() {
		defer wg.Done()
		users, uErr = queries.DouyinDB.GetUserInfos(userIDs)
	}()
	// follow
	go func() {
		defer wg.Done()
		curFollow, foErr = queries.DouyinDB.GetFollow(curID, userIDs)
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
	return c.JSON(http.StatusOK, utils.SuccessResponse(echo.Map{
		"next_time":  nextTime,
		"video_list": videos,
	}))
}
