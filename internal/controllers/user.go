package controllers

import (
	"garavel/internal/gates"
	"garavel/internal/libs"
	"garavel/internal/repositories"
	"garavel/internal/validators"
	"garavel/internal/views"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
}

func (ctr *User) Index(c *gin.Context) {
	list := repositories.User{}.GetList()
	repositories.Find(list)

	Success(c, http.StatusOK, views.List(list, views.User{}))
}

func (ctr *User) Create(c *gin.Context) {
	user, err := validators.Validate(c, validators.UserStore{}, repositories.User{}.GetModel())
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}

	user.Password, _ = libs.Hash(user.Password)
	repositories.Create(&user)

	Success(c, http.StatusCreated, views.Item(user, views.User{}))
}

func (ctr *User) Show(c *gin.Context) {
	if err := gates.Gate(c, gates.IsMeGate, c.Param("id")); err != nil {
		Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	user := repositories.User{}.GetModel()
	err := repositories.GetByID(&user, c.Param("id"))
	if err != nil {
		Error(c, http.StatusNotFound, "Record not found!")
		return
	}

	Success(c, http.StatusOK, views.Item(user, views.User{}))
}

func (ctr *User) Update(c *gin.Context) {
	if err := gates.Gate(c, gates.IsMeGate, c.Param("id")); err != nil {
		Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	user := repositories.User{}.GetModel()
	err := repositories.GetByID(&user, c.Param("id"))
	if err != nil {
		Error(c, http.StatusNotFound, "Record not found!")
		return
	}

	user.Password, _ = libs.Hash(user.Password)
	_, err = validators.Validate(c, validators.UserUpdate{}, &user)
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}

	err = repositories.Save(&user)
	if err != nil {
		Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	Success(c, http.StatusOK, views.Item(user, views.User{}))
}

func (ctr *User) Destroy(c *gin.Context) {
	if err := gates.Gate(c, gates.IsMeGate, c.Param("id")); err != nil {
		Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	user := repositories.User{}.GetModel()
	err := repositories.GetByID(&user, c.Param("id"))
	if err != nil {
		Error(c, http.StatusNotFound, "Record not found!")
		return
	}

	repositories.Delete(&user)

	Success(c, http.StatusNoContent, nil)
}
