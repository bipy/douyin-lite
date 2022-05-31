package controllers

import (
	"douyin-lite/app/queries"
	"douyin-lite/pkg/repository"
	"douyin-lite/pkg/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func Register(c echo.Context) error {
	params := c.QueryParams()
	username, password := params.Get("username"), params.Get("password")
	if username == "" || len(username) > repository.MaxUsernameLength {
		return c.JSON(http.StatusOK, utils.FailResponse("Illegal Username"))
	}
	if password == "" || len(password) > repository.MaxUserPasswordLength {
		return c.JSON(http.StatusOK, utils.FailResponse("Illegal Password"))
	}
	hash, err := utils.GeneratePassword(password)
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse(err.Error()))
	}
	id, err := queries.DouyinDB.CreateUser(username, hash)
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, utils.SuccessResponse(echo.Map{
		"user_id": id,
		"token":   "",
	}))
}

func Login(c echo.Context) error {
	params := c.QueryParams()
	username, password := params.Get("username"), params.Get("password")
	if username == "" || len(username) > repository.MaxUsernameLength {
		return c.JSON(http.StatusOK, utils.FailResponse("Illegal Username"))
	}
	if password == "" || len(password) > repository.MaxUserPasswordLength {
		return c.JSON(http.StatusOK, utils.FailResponse("Illegal Password"))
	}
	id, hash, err := queries.DouyinDB.GetHashByUsername(username)
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse(err.Error()))
	}
	if !utils.ComparePasswords(hash, password) {
		return c.JSON(http.StatusOK, utils.FailResponse("Username or Password is Incorrect"))
	}
	return c.JSON(http.StatusOK, utils.SuccessResponse(echo.Map{
		"user_id": id,
		"token":   "",
	}))
}

func GetUserInfo(c echo.Context) error {
	params := c.QueryParams()
	token := params.Get("token")
	if token == "" {
		return c.JSON(http.StatusOK, utils.FailResponse("Unauthorized"))
	}
	userID, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse("Illegal user_id"))
	}
	user, err := queries.DouyinDB.GetUserInfo(userID)
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, utils.SuccessResponse(echo.Map{
		"user": user,
	}))
}
