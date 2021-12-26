package server

import (
	"github.com/Ladence/golang_base_rest/internal/bl"
	"github.com/Ladence/golang_base_rest/internal/middleware"
	"github.com/gin-gonic/gin"
)

const (
	UserRootEndpoint = "/user"

	LoginEndpoint = "/login"
	RefreshTokenEndpoint = "/refresh_token"
	HelloEndpoint = "/hello"
)

func registerUserRoutes(gen *gin.Engine) error {
	authMw, err := middleware.NewAuthMiddleware()
	if err != nil {
		return err
	}
	err = authMw.MiddlewareInit()
	if err != nil {
		return err
	}

	gen.POST(LoginEndpoint, authMw.LoginHandler)

	auth := gen.Group(UserRootEndpoint)
	auth.Use(authMw.MiddlewareFunc())
	auth.GET(RefreshTokenEndpoint, authMw.RefreshHandler)
	auth.GET(HelloEndpoint, bl.Hello)
	return nil
}
