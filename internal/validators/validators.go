package validators

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
)

func Validate[R any, M any](c *gin.Context, rule R, model M) (*M, error) {
	if err := c.ShouldBindJSON(&rule); err != nil {
		return nil, err
	}

	if reflect.TypeOf(rule).Kind() != reflect.Struct {
		return nil, fmt.Errorf("rule is not a struct")
	}

	validate := validator.New()
	err := validate.Struct(rule)

	if err != nil {
		return nil, fmt.Errorf("validation failed: %v", err)
	}

	copier.Copy(&model, &rule)

	return &model, nil
}
