package factories

import "garavel/internal/models"

func Make[M any](class string) M {
	var result any
	switch class {
	case models.UserClass:
		result = User{}.Make()
	case models.CustomerClass:
		result = Customer{}.Make()
	}
	return result.(M)
}
