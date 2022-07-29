package main

import (
	"go_mvc/my-app/src/controller"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(os.Stdout, f)
	r := controller.GetRouter()
	r.Run(":8081")
}
