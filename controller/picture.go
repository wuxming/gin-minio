package controller

import (
	"gin-minio/service"
	"gin-minio/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

//上传文件(图片)
func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorMess(err.Error(), nil))
		return
	}
	result := service.Upload(file)
	c.JSON(result["status"].(int), result)
}
