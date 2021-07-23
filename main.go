package main

import (
	"gin-minio/router"
)

func main() {
	r := router.SetupRouter()
	err := r.Run(":8000")
	if err != nil {
		panic(err)
	}
}
