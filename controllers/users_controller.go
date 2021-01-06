package controllers

import (
	"net/http"

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

}

func SearchUser(c *gin.Context) {

}
