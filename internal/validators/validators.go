package validators

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Validate[R any](c *gin.Context, rule R) (*R, error) {
	if reflect.TypeOf(rule).Kind() != reflect.Struct {
		return nil, fmt.Errorf("rule is not a struct")
	}

	if err := c.ShouldBindJSON(&rule); err != nil {
		return nil, err
	}

	validate := validator.New()
	err := validate.Struct(rule)
	if err != nil {
		return nil, fmt.Errorf("validation failed: %v", err)
	}

	return &rule, nil
}
