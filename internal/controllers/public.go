package controllers

import (
	"garavel/internal/services"
	"garavel/internal/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Public struct {
}

func Health(c *gin.Context) {
	Success(c, http.StatusOK, "running...")
}

func Login(c *gin.Context) {
	request, err := validators.Validate(c, validators.UserStore{})
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := services.Auth{}.Login(request)
	if err != nil {
		Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	Success(c, http.StatusOK, gin.H{"token": token})
}
