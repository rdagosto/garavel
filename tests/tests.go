package tests

import (
	"garavel/internal/libs"
	"garavel/internal/models"
	"garavel/internal/routes"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUp() {
	routes.HandleRequest()
}

func TearDown() {
}

func GetRouter() *gin.Engine {
	return routes.R
}

func AddToken(req *http.Request) {
	jwtToken, _ := libs.GenerateJWT(123)
	req.Header.Set("Authorization", "Bearer "+jwtToken)
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
