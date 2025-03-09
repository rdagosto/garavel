package controllers

import (
	"garavel/internal/libs"
	"garavel/internal/models"
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
	user, err := validators.Validate(c, validators.UserStore{}, models.User{})
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var dbUser models.User
	if err := models.GetDB().Where("email = ?", user.Email).First(&dbUser).Error; err != nil {
		Error(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	if !libs.CheckHash(dbUser.Password, user.Password) {
		Error(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	token, err := libs.GenerateJWT(dbUser.ID)
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}

	Success(c, http.StatusOK, gin.H{"token": token})
}
