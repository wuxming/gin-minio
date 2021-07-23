package utils

import (
	"github.com/minio/minio-go"
	"log"
)

// 初使化 minio client对象。
func MinioClient() *minio.Client {
	endpoint := "42.193.99.32:9800"
	accessKeyID := "admin"
	secretAccessKey := "admin123456"
	useSSL := false
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}
	return minioClient
}
// 检查存储桶是否已经存在。
func BucketExists(minioClient *minio.Client,bucketName string) bool  {
	exists, err := minioClient.BucketExists(bucketName)
	if err != nil  {
		log.Println(err)
		return false
	}
	if  !exists {
		log.Println(bucketName,"不存在")
		return false
	}
	return true
}
