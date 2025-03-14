package controllers

import (
	"garavel/internal/models"
	"garavel/internal/services"
	"garavel/internal/validators"
	"garavel/internal/views"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Customer struct {
}

func (ctr *Customer) Index(c *gin.Context) {
	list := services.List[models.Customer]()
	Success(c, http.StatusOK, views.List(list, views.Customer{}))
}

func (ctr *Customer) Create(c *gin.Context) {
	request, err := validators.Validate(c, validators.CustomerStore{})
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}

	model, err := services.Create[models.Customer](request)
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}

	Success(c, http.StatusCreated, views.Item(model, views.Customer{}))
}

func (ctr *Customer) Show(c *gin.Context) {
	model, err := services.Show[models.Customer](c.Param("id"))
	if err != nil {
		Error(c, http.StatusNotFound, err.Error())
		return
	}

	Success(c, http.StatusOK, views.Item(model, views.Customer{}))
}

func (ctr *Customer) Update(c *gin.Context) {
	request, err := validators.Validate(c, validators.CustomerUpdate{})
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}

	model, err := services.Update[models.Customer](c.Param("id"), request)
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}

	Success(c, http.StatusOK, views.Item(model, views.Customer{}))
}

func (ctr *Customer) Destroy(c *gin.Context) {
	err := services.Delete[models.Customer](c.Param("id"))
	if err != nil {
		Error(c, http.StatusNotFound, "Record not found!")
		return
	}

	Success(c, http.StatusNoContent, nil)
}
