package controllers

import (
	"douyin-lite/pkg/configs"
	"douyin-lite/pkg/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func GetFile(c echo.Context) error {
	contentType := c.Param("type")
	id := c.Param("uuid")
	file, err := os.ReadFile(fmt.Sprintf("%s%s/%s", configs.FilePrefix, contentType, id))
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse("Get Static Resources Failed"))
	}
	if contentType == "cover" {
		return c.Blob(http.StatusOK, "image/jpeg", file)
	} else if contentType == "play" {
		return c.Blob(http.StatusOK, "video/mp4", file)
	}
	return c.NoContent(http.StatusOK)
}
