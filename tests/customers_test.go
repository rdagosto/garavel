package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"garavel/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	SetUp()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/customers", nil)
	AddToken(req)
	GetRouter().ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var customers []models.Customer
	json.Unmarshal(w.Body.Bytes(), &customers)
	// assert.Equal(t, 4, len(customers))
}

func TestShow(t *testing.T) {
	SetUp()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/customers/1", nil)
	AddToken(req)
	GetRouter().ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var customer models.Customer
	json.Unmarshal(w.Body.Bytes(), &customer)
	assert.Equal(t, 1, int(customer.ID))
}

func TestCreate(t *testing.T) {
	SetUp()
	defer TearDown()
	customer := models.Customer{Name: "John1", Email: "johndoe@example.com"}
	jsonData, _ := json.Marshal(customer)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/customers", bytes.NewBuffer(jsonData))
	AddToken(req)
	GetRouter().ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	var createdCustomer models.Customer
	json.Unmarshal(w.Body.Bytes(), &createdCustomer)
	assert.Equal(t, customer.Name, createdCustomer.Name)
	assert.Equal(t, customer.Email, createdCustomer.Email)
	assert.NotZero(t, createdCustomer.ID)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/customers/"+fmt.Sprint(createdCustomer.ID), nil)
	AddToken(req)
	GetRouter().ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)
	AssertDatabaseMissing(t, "customers", models.Customer{ID: createdCustomer.ID})
}

func TestUpdate(t *testing.T) {
	SetUp()
	defer TearDown()
	//TODO: seed to cresate
	customer := models.Customer{Name: "John2", Email: "johndoe@example.com"}
	jsonData, _ := json.Marshal(customer)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/customers", bytes.NewBuffer(jsonData))
	AddToken(req)
	req.Header.Set("Content-Type", "application/json")
	GetRouter().ServeHTTP(w, req)
	var createdCustomer models.Customer
	json.Unmarshal(w.Body.Bytes(), &createdCustomer)
	updatedCustomer := models.Customer{Name: "John Updated", Email: "johnupdated@example.com"}
	updatedJsonData, _ := json.Marshal(updatedCustomer)
	w = httptest.NewRecorder()

	req, _ = http.NewRequest("PUT", "/customers/"+fmt.Sprint(createdCustomer.ID), bytes.NewBuffer(updatedJsonData))
	AddToken(req)
	GetRouter().ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var updated models.Customer
	json.Unmarshal(w.Body.Bytes(), &updated)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, updatedCustomer.Name, updated.Name)
	assert.Equal(t, updatedCustomer.Email, updated.Email)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/customers/"+fmt.Sprint(createdCustomer.ID), nil)
	AddToken(req)
	GetRouter().ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)
	AssertDatabaseMissing(t, "customers", models.Customer{ID: createdCustomer.ID})
}

func TestDelete(t *testing.T) {
	SetUp()
	customer := models.Customer{Name: "John3", Email: "johndoe@example.com"}
	jsonData, _ := json.Marshal(customer)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/customers", bytes.NewBuffer(jsonData))
	AddToken(req)
	GetRouter().ServeHTTP(w, req)
	var createdCustomer models.Customer
	json.Unmarshal(w.Body.Bytes(), &createdCustomer)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/customers/"+fmt.Sprint(createdCustomer.ID), nil)
	AddToken(req)
	GetRouter().ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)
	AssertDatabaseMissing(t, "customers", models.Customer{ID: createdCustomer.ID})
}
