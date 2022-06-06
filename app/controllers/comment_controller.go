package controllers

import (
	"douyin-lite/app/models"
	"douyin-lite/app/queries"
	"douyin-lite/pkg/repository"
	"douyin-lite/pkg/utils"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func CommentAction(c echo.Context) error {
	curID, err := MustAuthorize(c)
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
		commentText := params.Get("comment_text")
		if commentText == "" || len(commentText) > repository.MaxVideoCommentTextLength {
			return c.JSON(http.StatusOK, utils.FailResponse("Illegal comment_text"))
		}

		wg := sync.WaitGroup{}
		wg.Add(2)

		var user *models.User
		var commentID int
		var cErr, uErr error

		go func() {
			defer wg.Done()
			commentID, cErr = queries.DouyinDB.CreateComment(curID, videoID, commentText)
		}()

		go func() {
			defer wg.Done()
			user, uErr = queries.DouyinDB.GetUserInfo(curID)
		}()

		wg.Wait()

		if cErr != nil || uErr != nil {
			return c.JSON(http.StatusOK, utils.FailResponse("Get Data Failed"))
		}

		comment, err := queries.DouyinDB.GetComment(commentID)
		if err != nil {
			return c.JSON(http.StatusOK, utils.FailResponse(err.Error()))
		}

		go func() {
			err := queries.DouyinDB.AddCommentCount(1, videoID)
			if err != nil {
				log.Fatal(err.Error())
			}
		}()

		comment.CreateDate = utils.UnixTimeStampToCSTMMdd(comment.CreateTime)
		comment.Author = user

		return c.JSON(http.StatusOK, utils.SuccessResponse(echo.Map{
			"comment": comment,
		}))
	} else if actionType == 2 {
		commentID, err := strconv.Atoi(params.Get("comment_id"))
		if err != nil {
			return c.JSON(http.StatusOK, utils.FailResponse("Illegal comment_id"))
		}

		effected, err := queries.DouyinDB.DeleteComment(commentID, curID)
		if err != nil {
			return c.JSON(http.StatusOK, utils.FailResponse(err.Error()))
		}

		if effected == 0 {
			return c.JSON(http.StatusOK, utils.FailResponse("Delete Failed"))
		} else {
			go func() {
				err := queries.DouyinDB.AddCommentCount(-1, videoID)
				if err != nil {
					log.Fatal(err.Error())
				}
			}()
		}
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse(echo.Map{
		"comment": nil,
	}))
}

func GetCommentList(c echo.Context) error {
	curID, err := MustAuthorize(c)
	if err != nil {
		return err
	}

	videoID, err := strconv.Atoi(c.QueryParam("video_id"))
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse("Illegal video_id"))
	}

	comments, err := queries.DouyinDB.GetComments(videoID)
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse(err.Error()))
	}

	if len(comments) > 0 {
		userIDs := make([]int, len(comments))
		for i, v := range comments {
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
		for i := range comments {
			comments[i].Author = userMap[comments[i].AuthorId]
			comments[i].CreateDate = utils.UnixTimeStampToCSTMMdd(comments[i].CreateTime)
		}
	}
	return c.JSON(http.StatusOK, utils.SuccessResponse(echo.Map{
		"comment_list": comments,
	}))
}
