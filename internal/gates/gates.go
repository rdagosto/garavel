package gates

import (
	"errors"
	"garavel/internal/models"

	"github.com/gin-gonic/gin"
)

const (
	IsMeGate = "user"
)

func Gate(c *gin.Context, gate string, params ...string) error {
	user, exists := c.Get("user")
	if !exists {
		return errors.New("user not found")
	}

	switch gate {
	case IsMeGate:
		return IsMe(c, user.(models.User), params)
	default:
		return errors.New("Gate not found")
	}
}
