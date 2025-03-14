package controllers

import (
	"garavel/internal/repositories"
	"garavel/internal/validators"
	"garavel/internal/views"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Customer struct {
}

func (ctr *Customer) Index(c *gin.Context) {
	list := repositories.Customer{}.GetList()
	repositories.Find(list)
	Success(c, http.StatusOK, views.List(list, views.Customer{}))
}

func (ctr *Customer) Create(c *gin.Context) {
	customer, err := validators.Validate(c, validators.CustomerStore{}, repositories.Customer{}.GetModel())
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}
	repositories.Create(&customer)
	Success(c, http.StatusCreated, views.Item(customer, views.Customer{}))
}

func (ctr *Customer) Show(c *gin.Context) {
	customer := repositories.Customer{}.GetModel()
	err := repositories.GetByID(&customer, c.Param("id"))
	if err != nil {
		Error(c, http.StatusNotFound, "Record not found!")
		return
	}
	Success(c, http.StatusOK, views.Item(customer, views.Customer{}))
}

func (ctr *Customer) Update(c *gin.Context) {
	customer := repositories.Customer{}.GetModel()
	err := repositories.GetByID(&customer, c.Param("id"))
	if err != nil {
		Error(c, http.StatusNotFound, "Record not found!")
		return
	}

	_, err = validators.Validate(c, validators.CustomerUpdate{}, &customer)
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}

	err = repositories.Save(&customer)
	if err != nil {
		Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	Success(c, http.StatusOK, views.Item(customer, views.Customer{}))
}

func (ctr *Customer) Destroy(c *gin.Context) {
	customer := repositories.Customer{}.GetModel()
	err := repositories.GetByID(&customer, c.Param("id"))
	if err != nil {
		Error(c, http.StatusNotFound, "Record not found!")
		return
	}
	repositories.Delete(&customer)
	Success(c, http.StatusNoContent, nil)
}
