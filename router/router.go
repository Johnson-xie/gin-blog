package router

import (
	"github.com/gin-gonic/gin"
	"github.com/linehk/gin-blog/config"
	"github.com/linehk/gin-blog/router/api/v1"
)

func Init() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(config.Server.RunMode)
	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/articles", v1.Articles)
	}
	return r
}
