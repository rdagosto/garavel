package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	Index(c *gin.Context)
	Create(c *gin.Context)
	Show(c *gin.Context)
	Update(c *gin.Context)
	Destroy(c *gin.Context)
}

func Factory(ctr string) Controller {
	switch ctr {
	case "customers":
		return &Customer{}
	case "users":
		return &User{}
	default:
		log.Fatalf("❌ Unsupported controller type: %s", ctr)
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
