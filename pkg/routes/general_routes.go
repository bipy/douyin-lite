package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GeneralRoutes(a *echo.Echo) {
	a.GET("/generate_204", func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusNoContent)
	})
}
