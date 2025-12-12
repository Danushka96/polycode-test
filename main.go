package main

import (
	"fmt"
	"pos-app/controllers"

	runtime "github.com/cloudimpl/polycode-runtime/go"
	"github.com/gin-gonic/gin"
	_ "pos-app/.polycode"
)

func main() {
	r := gin.Default()

	r.GET("/hello", controllers.HelloWithoutCache)
	r.GET("/hello-cache", controllers.HelloWithCache)
	r.POST("/hello-post", controllers.HelloPOST)

	err := runtime.Start(runtime.WithHttpHandler(r))
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to start runtime: %w", err))
		return
	}
}
