package service

import (
	"gin1/utils"
	"github.com/minio/minio-go"
	"mime/multipart"
)

func Upload(file *multipart.FileHeader)map[string]interface{} {

	src, err := file.Open()//不懂
	if err != nil {
		return utils.ErrorMess(err.Error(), nil)
	}
	defer src.Close()
	minioClient := utils.MinioClient()
	bucketName := "test"
	fileName :="wxm"+file.Filename//要前缀
	_,err = minioClient.PutObject(bucketName, fileName, src, -1,minio.PutObjectOptions{ContentType: file.Header.Get("Content-Type")})
	if err != nil {
		return utils.ErrorMess(err.Error(), nil)
	}
	pictureUrl := "http://42.193.99.32:9800/"+bucketName+"/"+fileName
	return utils.SuccessMess("图片地址", pictureUrl)
}