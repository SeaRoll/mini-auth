package service

import (
	"net/http"

	"github.com/SeaRoll/mini-auth/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/google/uuid"
)\

type googleOAuthService struct {
	config config.Config
}

func NewGoogleOAuthService(config config.Config) OAuthServiceStrategy {
	return &googleOAuthService{
		config: config,
	}
}

// Redirect implements OAuthServiceStrategy.
func (g *googleOAuthService) Redirect(c echo.Context) error {
	stateCode := uuid.New().String()
	url := g.config.Auths.Google.OAuthConfig.AuthCodeURL(stateCode)
	return c.Redirect(http.StatusFound, url)
}

// PerformCallback implements OAuthServiceStrategy.
func (g *googleOAuthService) PerformCallback(c echo.Context) error {
	panic("unimplemented")
}
