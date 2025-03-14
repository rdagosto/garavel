package factories

import (
	"garavel/internal/models"
	"reflect"
)

func Make[M any](class string, attrs map[string]any) M {
	var entity any
	switch class {
	case models.UserClass:
		entity = User{}.Make()
	case models.CustomerClass:
		entity = Customer{}.Make()
	}
	entity = merge(entity, attrs)
	return entity.(M)
}

func merge[T any](instance T, attrs map[string]interface{}) T {
	structValue := reflect.ValueOf(&instance).Elem()

	for key, value := range attrs {
		field := structValue.FieldByName(key)
		if field.IsValid() && field.CanSet() {
			field.Set(reflect.ValueOf(value))
		}
	}
	return instance
}
