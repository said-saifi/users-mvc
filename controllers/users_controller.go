package controllers

import (
	"net/http"
	"strconv"

	"github.com/said-saifi/users-mvc/utils/errors"

	"github.com/said-saifi/users-mvc/services"

	"github.com/gin-gonic/gin"
	"github.com/said-saifi/users-mvc/domain/users"
)

func CreateUser(c *gin.Context) {
	var user users.User

	// this reads the body and unmarshal the json
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	res, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, res)
}

func GetUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("user id should be an integer")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, restErr := services.GetUser(id)
	if err != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func SearchUser(c *gin.Context) {

}
