package controllers

import (
	"garavel/internal/gates"
	"garavel/internal/models"
	"garavel/internal/services"
	"garavel/internal/validators"
	"garavel/internal/views"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
}

func (ctr *User) Index(c *gin.Context) {
	list := services.List[models.User]()
	Success(c, http.StatusOK, views.List(list, views.User{}))
}

func (ctr *User) Create(c *gin.Context) {
	request, err := validators.Validate(c, validators.UserStore{})
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}

	model, err := services.User{}.Create(request)
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}

	Success(c, http.StatusCreated, views.Item(model, views.User{}))
}

func (ctr *User) Show(c *gin.Context) {
	if err := gates.Gate(c, gates.IsMeGate, c.Param("id")); err != nil {
		Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	model, err := services.Show[models.User](c.Param("id"))
	if err != nil {
		Error(c, http.StatusNotFound, err.Error())
		return
	}

	Success(c, http.StatusOK, views.Item(model, views.User{}))
}

func (ctr *User) Update(c *gin.Context) {
	if err := gates.Gate(c, gates.IsMeGate, c.Param("id")); err != nil {
		Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	request, err := validators.Validate(c, validators.UserUpdate{})
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}

	model, err := services.User{}.Update(c.Param("id"), request)
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}

	Success(c, http.StatusOK, views.Item(model, views.User{}))
}

func (ctr *User) Destroy(c *gin.Context) {
	if err := gates.Gate(c, gates.IsMeGate, c.Param("id")); err != nil {
		Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	err := services.Delete[models.User](c.Param("id"))
	if err != nil {
		Error(c, http.StatusNotFound, "Record not found!")
		return
	}

	Success(c, http.StatusNoContent, nil)
}
