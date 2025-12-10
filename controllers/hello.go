package controllers

import (
	"pos-app/models"

	"github.com/cloudimpl/polycode-runtime/go/apicontext"
	"github.com/cloudimpl/polycode-runtime/go/sdk"
	"github.com/gin-gonic/gin"
)

func HelloWithCache(c *gin.Context) {
	var req map[string]interface{}
	req["name"] = c.Query("name")

	ctx, err := apicontext.FromContext(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	greetingService := ctx.Service("greeting-service").Get()
	reply, err := greetingService.RequestReply(sdk.TaskOptions{}, "Hello", req)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	var res models.HelloResponse
	err = reply.Get(&res)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, res)
}

func HelloWithoutCache(c *gin.Context) {
	var req map[string]interface{}
	req["name"] = c.Query("name")

	ctx, err := apicontext.FromContext(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	greetingService := ctx.Service("greeting-service").Get()
	reply, err := greetingService.RequestReply(sdk.TaskOptions{}, "Hello", req)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	var res models.HelloResponse
	err = reply.Get(&res)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Header("Cache-Control", "public, max-age=300")
	c.JSON(200, res)
}

func HelloPOST(c *gin.Context) {
	var req map[string]interface{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, err := apicontext.FromContext(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	greetingService := ctx.Service("greeting-service").Get()
	reply, err := greetingService.RequestReply(sdk.TaskOptions{}, "Hello", req)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	var res models.HelloResponse
	err = reply.Get(&res)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Header("Cache-Control", "public, max-age=300")
	c.JSON(200, res)
}
