package controllers

import (
	"garavel/internal/models"
	"garavel/internal/validators"
	"garavel/internal/views"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Customer struct {
}

func (ctr *Customer) Index(c *gin.Context) {
	var customers []models.Customer
	models.GetDB(models.Factory("customer")).Find(&customers)
	Success(c, http.StatusOK, views.List(customers, views.Customer{}))
}

func (ctr *Customer) Create(c *gin.Context) {
	customer, err := validators.Validate(c, validators.CustomerStore{}, models.Customer{})
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}
	models.GetDB(models.Factory("customer")).Create(&customer)
	Success(c, http.StatusCreated, views.Item(customer, views.Customer{}))
}

func (ctr *Customer) Show(c *gin.Context) {
	var customer models.Customer
	if err := models.GetDB(models.Factory("customer")).First(&customer, c.Param("id")).Error; err != nil {
		Error(c, http.StatusNotFound, "Record not found!")
		return
	}
	Success(c, http.StatusOK, views.Item(customer, views.Customer{}))
}

func (ctr *Customer) Update(c *gin.Context) {
	var customer models.Customer
	if err := models.GetDB(models.Factory("customer")).First(&customer, c.Param("id")).Error; err != nil {
		Error(c, http.StatusNotFound, "Record not found!")
		return
	}

	_, err := validators.Validate(c, validators.CustomerUpdate{}, &customer)
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := models.GetDB(models.Factory("customer")).Save(&customer).Error; err != nil {
		Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	Success(c, http.StatusOK, views.Item(customer, views.Customer{}))
}

func (ctr *Customer) Destroy(c *gin.Context) {
	var customer models.Customer
	if err := models.GetDB(models.Factory("customer")).First(&customer, c.Param("id")).Error; err != nil {
		Error(c, http.StatusNotFound, "Record not found!")
		return
	}
	models.GetDB(models.Factory("customer")).Delete(&customer)
	Success(c, http.StatusNoContent, nil)
}
