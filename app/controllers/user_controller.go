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

	token, err := utils.GenerateToken(id)
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse(echo.Map{
		"user_id": id,
		"token":   token,
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

	token, err := utils.GenerateToken(id)
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse(echo.Map{
		"user_id": id,
		"token":   token,
	}))
}

func GetUserInfo(c echo.Context) error {
	curID, err := MustAuthorize(c)
	if err != nil {
		return err
	}

	userID, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse("Illegal user_id"))
	}

	var user *models.User
	isFollow := false
	var uErr, foErr error

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		user, uErr = queries.DouyinDB.GetUserInfo(userID)
	}()

	go func() {
		defer wg.Done()
		isFollow, foErr = queries.DouyinDB.IsFollow(curID, userID)
	}()

	wg.Wait()

	if uErr != nil {
		return c.JSON(http.StatusOK, utils.FailResponse(uErr.Error()))
	}

	if foErr != nil {
		return c.JSON(http.StatusOK, utils.FailResponse(foErr.Error()))
	}

	user.IsFollow = isFollow

	return c.JSON(http.StatusOK, utils.SuccessResponse(echo.Map{
		"user": user,
	}))
}
