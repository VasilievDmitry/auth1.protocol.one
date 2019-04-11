package route

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func InitHealth(cfg Config) error {
	cfg.Echo.GET("/health", health)

	return nil
}

func health(ctx echo.Context) error {
	return ctx.HTML(http.StatusNoContent, ``)
}