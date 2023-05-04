package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/facundocenteno/bookstore_users-api/domain/users"
	"github.com/facundocenteno/bookstore_users-api/services"
	"github.com/facundocenteno/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	res, restErr := services.CreateUser(user)
	if restErr != nil {
		restErr := errors.NewNotFoundError("error at creating new user")
		c.JSON(restErr.Status, restErr)
		return
	}

	fmt.Println(user)
	c.JSON(http.StatusCreated, res)
}

func FindUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		restErr := errors.NewBadRequestError("user should be a number")
		c.JSON(restErr.Status, restErr)
	}

	res, getErr := services.FindUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, nil)
}
