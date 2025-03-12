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
	"syreclabs.com/go/faker"
)

func TestUserIndex(t *testing.T) {
	SetUp()
	defer TearDown()
	Create[models.User](models.UserClass, "user", map[string]any{})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	AddToken(req, GetEntity[models.User]("user").ID)
	GetRouter().ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserShow(t *testing.T) {
	SetUp()
	defer TearDown()
	Create[models.User](models.UserClass, "user", map[string]any{})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/users/%d", GetEntity[models.User]("user").ID), nil)
	AddToken(req, GetEntity[models.User]("user").ID)
	GetRouter().ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var user models.User
	json.Unmarshal(w.Body.Bytes(), &user)
	assert.Equal(t, GetEntity[models.User]("user").ID, user.ID)
}

func TestUserCreate(t *testing.T) {
	SetUp()
	defer TearDown()
	Create[models.User](models.UserClass, "user", map[string]any{})

	payload := map[string]any{
		"email":    faker.Internet().Email(),
		"password": "securePass123",
	}
	jsonData, _ := json.Marshal(payload)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonData))
	AddToken(req, GetEntity[models.User]("user").ID)
	GetRouter().ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	var createdUser models.User
	json.Unmarshal(w.Body.Bytes(), &createdUser)
	assert.Equal(t, payload["email"], createdUser.Email)
	assert.NotZero(t, createdUser.ID)
}

func TestUserUpdate(t *testing.T) {
	SetUp()
	defer TearDown()
	Create[models.User](models.UserClass, "user", map[string]any{})

	payload := map[string]any{
		"password": "securePass123",
	}
	jsonData, _ := json.Marshal(payload)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/users/%d", GetEntity[models.User]("user").ID), bytes.NewBuffer(jsonData))
	AddToken(req, GetEntity[models.User]("user").ID)
	GetRouter().ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserDelete(t *testing.T) {
	SetUp()
	defer TearDown()
	Create[models.User](models.UserClass, "user", map[string]any{})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/users/%d", GetEntity[models.User]("user").ID), nil)
	AddToken(req, GetEntity[models.User]("user").ID)
	GetRouter().ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)
	AssertDatabaseMissing(t, "users", models.User{ID: GetEntity[models.User]("user").ID})
}
