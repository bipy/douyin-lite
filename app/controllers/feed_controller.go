package controllers

import (
	"douyin-lite/app/models"
	"douyin-lite/pkg/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetHello(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello World!")
}

func GetFeed(c echo.Context) error {
	params := c.QueryParams()
	userId := params.Get("user_id")
	if userId == "" {
		return c.JSON(http.StatusOK, utils.FailResponse("user_id is empty"))
	}
	return c.JSON(http.StatusOK, utils.SuccessResponse(echo.Map{
		"user": models.User{
			Id:            0,
			Name:          "H",
			FollowCount:   0,
			FollowerCount: 0,
			CreateTime:    0,
		},
	}))
}
