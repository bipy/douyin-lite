package controllers

import (
	"douyin-lite/app/models"
	"douyin-lite/app/queries"
	"douyin-lite/pkg/utils"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

func RelationAction(c echo.Context) error {
	curID, err := MustAuthorize(c)
	if err != nil {
		return err
	}
	params := c.QueryParams()
	toUserID, err := strconv.Atoi(params.Get("to_user_id"))
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse("Illegal to_user_id"))
	}

	if curID == toUserID {
		return c.JSON(http.StatusOK, utils.FailResponse("不能关注自己噢"))
	}

	actionType, err := strconv.Atoi(params.Get("action_type"))
	if err != nil || actionType != 1 && actionType != 2 {
		return c.JSON(http.StatusOK, utils.FailResponse("Illegal action_type"))
	}

	if actionType == 1 {
		err := queries.DouyinDB.DoFollow(curID, toUserID)
		if err != nil {
			return c.JSON(http.StatusOK, utils.FailResponse(err.Error()))
		}
		go func() {
			err := queries.DouyinDB.AddFollowCount(1, curID)
			if err != nil {
				log.Fatal(err.Error())
			}
		}()
		go func() {
			err := queries.DouyinDB.AddFollowerCount(1, toUserID)
			if err != nil {
				log.Fatal(err.Error())
			}
		}()
	} else if actionType == 2 {
		err := queries.DouyinDB.CancelFollow(curID, toUserID)
		if err != nil {
			return c.JSON(http.StatusOK, utils.FailResponse(err.Error()))
		}
		go func() {
			err := queries.DouyinDB.AddFollowCount(-1, curID)
			if err != nil {
				log.Fatal(err.Error())
			}
		}()
		go func() {
			err := queries.DouyinDB.AddFollowerCount(-1, toUserID)
			if err != nil {
				log.Fatal(err.Error())
			}
		}()
	}
	return c.JSON(http.StatusOK, utils.SuccessResponse(echo.Map{}))
}

func GetFollowList(c echo.Context) error {
	curID, err := MustAuthorize(c)
	if err != nil {
		return err
	}

	userID, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse("Illegal user_id"))
	}

	users, err := queries.DouyinDB.GetFollowList(userID)
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse(err.Error()))
	}

	if curID != userID && len(users) > 0 {
		userIDs := make([]int, len(users))
		for i := range userIDs {
			userIDs[i] = users[i].Id
		}

		curFollow, err := queries.DouyinDB.GetFollows(curID, userIDs)
		if err != nil {
			return c.JSON(http.StatusOK, utils.FailResponse(err.Error()))
		}

		userMap := map[int]*models.User{}
		for i := range users {
			userMap[users[i].Id] = &users[i]
		}

		for _, f := range curFollow {
			userMap[f].IsFollow = true
		}
	} else {
		for i := range users {
			users[i].IsFollow = true
		}
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse(echo.Map{
		"user_list": users,
	}))
}

func GetFollowerList(c echo.Context) error {
	curID, err := MustAuthorize(c)
	if err != nil {
		return err
	}

	userID, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse("Illegal user_id"))
	}

	users, err := queries.DouyinDB.GetFollowerList(userID)
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse(err.Error()))
	}

	if len(users) > 0 {
		userIDs := make([]int, len(users))
		for i := range userIDs {
			userIDs[i] = users[i].Id
		}

		curFollow, err := queries.DouyinDB.GetFollows(curID, userIDs)
		if err != nil {
			return c.JSON(http.StatusOK, utils.FailResponse(err.Error()))
		}

		userMap := map[int]*models.User{}
		for i := range users {
			userMap[users[i].Id] = &users[i]
		}

		for _, f := range curFollow {
			userMap[f].IsFollow = true
		}
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse(echo.Map{
		"user_list": users,
	}))
}
