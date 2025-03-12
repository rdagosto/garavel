package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"garavel/internal/factories"
	"garavel/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomerIndex(t *testing.T) {
	SetUp()
	defer TearDown()
	Create[models.User](models.UserClass, "user", map[string]any{})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/customers", nil)
	AddToken(req, GetEntity[models.User]("user").ID)
	GetRouter().ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCustomerShow(t *testing.T) {
	SetUp()
	defer TearDown()
	Create[models.User](models.UserClass, "user", map[string]any{})
	Create[models.Customer](models.CustomerClass, "customer", map[string]any{})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/customers/%d", GetEntity[models.Customer]("customer").ID), nil)
	AddToken(req, GetEntity[models.User]("user").ID)
	GetRouter().ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var customer models.Customer
	json.Unmarshal(w.Body.Bytes(), &customer)
	assert.Equal(t, GetEntity[models.Customer]("customer").ID, customer.ID)
}

func TestCustomerCreate(t *testing.T) {
	SetUp()
	defer TearDown()
	Create[models.User](models.UserClass, "user", map[string]any{})

	customer := factories.Make[models.Customer](models.CustomerClass, map[string]any{})
	jsonData, _ := json.Marshal(customer)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/customers", bytes.NewBuffer(jsonData))
	AddToken(req, GetEntity[models.User]("user").ID)
	GetRouter().ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	var createdCustomer models.Customer
	json.Unmarshal(w.Body.Bytes(), &createdCustomer)
	assert.Equal(t, customer.Name, createdCustomer.Name)
	assert.Equal(t, customer.Email, createdCustomer.Email)
	assert.NotZero(t, createdCustomer.ID)
}

func TestCustomerUpdate(t *testing.T) {
	SetUp()
	defer TearDown()
	Create[models.User](models.UserClass, "user", map[string]any{})
	Create[models.Customer](models.CustomerClass, "customer", map[string]any{})

	updatedCustomer := factories.Make[models.Customer](models.CustomerClass, map[string]any{})
	updatedJsonData, _ := json.Marshal(updatedCustomer)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/customers/%d", GetEntity[models.Customer]("customer").ID), bytes.NewBuffer(updatedJsonData))
	AddToken(req, GetEntity[models.User]("user").ID)
	GetRouter().ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var updated models.Customer
	json.Unmarshal(w.Body.Bytes(), &updated)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, updatedCustomer.Name, updated.Name)
	assert.Equal(t, updatedCustomer.Email, updated.Email)
}

func TestCustomerDelete(t *testing.T) {
	SetUp()
	defer TearDown()
	Create[models.User](models.UserClass, "user", map[string]any{})
	Create[models.Customer](models.CustomerClass, "customer", map[string]any{})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/customers/%d", GetEntity[models.Customer]("customer").ID), nil)
	AddToken(req, GetEntity[models.User]("user").ID)
	GetRouter().ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)
	AssertDatabaseMissing(t, "customers", models.Customer{ID: GetEntity[models.User]("user").ID})
}
