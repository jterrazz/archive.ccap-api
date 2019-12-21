package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jterrazz/ccap.live-api/server"
	"github.com/jterrazz/ccap.live-api/utils"
)

func main() {
	cfg := utils.ReadConfig()
	router := gin.Default()

	router.POST("/graphql", server.GraphqlHandler())
	router.GET("/playground", server.PlaygroundHandler())

	_ = router.Run(cfg.Server.Host + ":" + cfg.Server.Port)
}
