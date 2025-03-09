package controllers

import (
	"garavel/internal/gates"
	"garavel/internal/libs"
	"garavel/internal/models"
	"garavel/internal/validators"
	"garavel/internal/views"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
}

func (ctr *User) Index(c *gin.Context) {
	var users []models.User
	models.GetDB().Find(&users)
	Success(c, http.StatusOK, views.List(users, views.User{}))
}

func (ctr *User) Create(c *gin.Context) {
	user, err := validators.Validate(c, validators.UserStore{}, models.User{})
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}
	user.Password, _ = libs.Hash(user.Password)
	models.GetDB().Create(&user)
	Success(c, http.StatusCreated, views.Item(user, views.User{}))
}

func (ctr *User) Show(c *gin.Context) {
	if err := gates.Gate(c, "isMe", c.Param("id")); err != nil {
		Error(c, http.StatusUnauthorized, err.Error())
		return
	}
	var user models.User
	if err := models.GetDB().First(&user, c.Param("id")).Error; err != nil {
		Error(c, http.StatusNotFound, "Record not found!")
		return
	}
	Success(c, http.StatusOK, views.Item(user, views.User{}))
}

func (ctr *User) Update(c *gin.Context) {
	if err := gates.Gate(c, "isMe", c.Param("id")); err != nil {
		Error(c, http.StatusUnauthorized, err.Error())
		return
	}
	var user models.User
	if err := models.GetDB().First(&user, c.Param("id")).Error; err != nil {
		Error(c, http.StatusNotFound, "Record not found!")
		return
	}

	_, err := validators.Validate(c, validators.UserUpdate{}, &user)
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}

	user.Password, _ = libs.Hash(user.Password)
	if err := models.GetDB().Save(&user).Error; err != nil {
		Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	Success(c, http.StatusOK, views.Item(user, views.User{}))
}

func (ctr *User) Destroy(c *gin.Context) {
	if err := gates.Gate(c, "isMe", c.Param("id")); err != nil {
		Error(c, http.StatusUnauthorized, err.Error())
		return
	}
	var user models.User
	if err := models.GetDB().First(&user, c.Param("id")).Error; err != nil {
		Error(c, http.StatusNotFound, "Record not found!")
		return
	}
	models.GetDB().Delete(&user)
	Success(c, http.StatusNoContent, nil)
}
