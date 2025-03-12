package factories

import (
	"garavel/internal/models"
	"reflect"
)

func Make[M any](class string, attrs map[string]any) M {
	var result any
	switch class {
	case models.UserClass:
		result = User{}.Make()
	case models.CustomerClass:
		result = Customer{}.Make()
	}
	applyOverrides(&result, attrs)
	return result.(M)
}

func applyOverrides(obj any, attrs map[string]any) {
	v := reflect.ValueOf(obj).Elem()
	for key, value := range attrs {
		field := v.FieldByName(key)
		if field.IsValid() && field.CanSet() {
			field.Set(reflect.ValueOf(value))
		}
	}
}
