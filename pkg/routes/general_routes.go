package routes

import (
	"douyin-lite/pkg/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func GeneralRoutes(a *echo.Echo) {
	a.GET("/generate_204", func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusNoContent)
	})

	a.GET("/sample-video", func(c echo.Context) error {
		video, err := os.ReadFile("sample.mp4")
		if err != nil {
			return c.JSON(http.StatusOK, utils.FailResponse("Get Sample Video Failed"))
		}
		return c.Blob(http.StatusOK, "video/mp4", video)
	})

	a.GET("/sample-cover", func(c echo.Context) error {
		cover, err := os.ReadFile("sample.jpg")
		if err != nil {
			return c.JSON(http.StatusOK, utils.FailResponse("Get Sample Cover Failed"))
		}
		return c.Blob(http.StatusOK, "image/jpeg", cover)
	})
}
