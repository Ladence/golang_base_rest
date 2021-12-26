package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Config struct {
	Host string `json:"host"`
	Port uint16 `json:"port"`
}

func Run(port int) error {
	gen := gin.New()
	addGinMiddlewares(gen)

	if err := registerUserRoutes(gen); err != nil {
		return err
	}
	return http.ListenAndServe(":"+strconv.Itoa(port), gen)
}

func addGinMiddlewares(gen *gin.Engine) {
	gen.Use(gin.Logger())
	gen.Use(gin.Recovery())
}