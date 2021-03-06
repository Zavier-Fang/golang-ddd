package http

import "github.com/gin-gonic/gin"

var Engine *gin.Engine

func init() {
	// gin.SetMode(gin.ReleaseMode)
	Engine = gin.Default()
}

