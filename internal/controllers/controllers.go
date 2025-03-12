package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
)

const (
	UserClass     = "users"
	CustomerClass = "customers"
)

type Controller interface {
	Index(c *gin.Context)
	Create(c *gin.Context)
	Show(c *gin.Context)
	Update(c *gin.Context)
	Destroy(c *gin.Context)
}

func Make(class string) Controller {
	switch class {
	case CustomerClass:
		return &Customer{}
	case UserClass:
		return &User{}
	default:
		log.Fatalf("❌ Unsupported controller type: %s", class)
		return nil
	}
}

func Success(c *gin.Context, status int, results interface{}) {
	c.JSON(status, results)
}

func Error(c *gin.Context, status int, msg string) {
	c.JSON(status, gin.H{"error": msg})
	c.Abort()
}
