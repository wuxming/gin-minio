package router

import (
	"gin1/controller"
	"gin1/utils"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.Use(utils.CORS(utils.Options{Origin: "*"})) //跨域
	Picture := r.Group("/picture")
	{
		Picture.POST("/upload", controller.Upload)
	}
	return r
}
