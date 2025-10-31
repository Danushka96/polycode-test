package main

import (
	"fmt"
	runtime "github.com/cloudimpl/polycode-runtime/go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "pos-app/.polycode"
	"pos-app/controllers"
)

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	r.POST("/hello", controllers.Hello)

	err := runtime.Start(runtime.WithHttpHandler(r))
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to start runtime: %w", err))
		return
	}
}
