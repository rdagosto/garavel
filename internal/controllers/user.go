package controllers

import (
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
	models.GetDB(models.Factory("user")).Find(&users)
	Success(c, http.StatusOK, views.List(users, views.User{}))
}

func (ctr *User) Create(c *gin.Context) {
	user, err := validators.Validate(c, validators.UserStore{}, models.User{})
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}
	user.Password, _ = libs.Hash(user.Password)
	models.GetDB(models.Factory("user")).Create(&user)
	Success(c, http.StatusCreated, views.Item(user, views.User{}))
}

func (ctr *User) Show(c *gin.Context) {
	var user models.User
	if err := models.GetDB(models.Factory("user")).First(&user, c.Param("id")).Error; err != nil {
		Error(c, http.StatusNotFound, "Record not found!")
		return
	}
	Success(c, http.StatusOK, views.Item(user, views.User{}))
}

func (ctr *User) Update(c *gin.Context) {
	var user models.User
	if err := models.GetDB(models.Factory("user")).First(&user, c.Param("id")).Error; err != nil {
		Error(c, http.StatusNotFound, "Record not found!")
		return
	}

	_, err := validators.Validate(c, validators.UserUpdate{}, &user)
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}

	user.Password, _ = libs.Hash(user.Password)
	if err := models.GetDB(models.Factory("user")).Save(&user).Error; err != nil {
		Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	Success(c, http.StatusOK, views.Item(user, views.User{}))
}

func (ctr *User) Destroy(c *gin.Context) {
	var user models.User
	if err := models.GetDB(models.Factory("user")).First(&user, c.Param("id")).Error; err != nil {
		Error(c, http.StatusNotFound, "Record not found!")
		return
	}
	models.GetDB(models.Factory("user")).Delete(&user)
	Success(c, http.StatusNoContent, nil)
}
