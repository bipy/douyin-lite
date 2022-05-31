package controllers

import (
	"douyin-lite/pkg/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Authorize(c echo.Context) (int, error) {
	token := c.QueryParam("token")
	userID, err := utils.Verify(token)
	if err != nil {
		return 0, c.JSON(http.StatusOK, utils.FailResponse("Unauthorized"))
	}
	return userID, nil
}
