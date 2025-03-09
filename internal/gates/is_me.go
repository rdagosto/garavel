package gates

import (
	"errors"
	"garavel/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IsMe(c *gin.Context, user models.User, params []string) error {
	ID, _ := strconv.ParseUint(params[0], 10, 0)
	if uint(ID) != user.ID {
		return errors.New("you are not allowed to access this resource")
	}
	return nil
}
