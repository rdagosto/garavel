package tests

import (
	"garavel/internal/factories"
	"garavel/internal/libs"
	"garavel/internal/models"
	"garavel/internal/routes"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var entities = make(map[string]any)

func GetEntity[M any](name string) M {
	value, exists := entities[name]
	if !exists {
		var empty M
		return empty
	}
	return value.(M)
}

func SetUp() {
	routes.HandleRequest()
}

func TearDown() {
}

func GetRouter() *gin.Engine {
	return routes.R
}

func AddToken(req *http.Request, userID uint) {
	jwtToken, _ := libs.GenerateJWT(userID)
	req.Header.Set("Authorization", "Bearer "+jwtToken)
}

func Create[M any](class string, name string, attrs map[string]any) {
	model := factories.Make[M](class, attrs)
	models.GetDB().Create(&model)
	entities[name] = model
}

func AssertDatabaseHas(t *testing.T, table string, condition interface{}) {
	var count int64
	err := models.GetDB().Table(table).Where(condition).Count(&count).Error
	if err != nil {
		t.Fatalf("Error checking database: %v", err)
	}
	assert.Greater(t, count, int64(0), "Database does not contain the expected record")
}

func AssertDatabaseMissing(t *testing.T, table string, condition interface{}) {
	var count int64
	err := models.GetDB().Table(table).Where(condition).Count(&count).Error
	if err != nil {
		t.Fatalf("Error checking database: %v", err)
	}
	assert.Equal(t, count, int64(0), "Database does not contain the expected record")
}
