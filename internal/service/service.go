package service

import (
	"github.com/labstack/echo/v4"
)

type OAuthServiceStrategy interface {
	Redirect(c echo.Context) error
	PerformCallback(c echo.Context) error
}
