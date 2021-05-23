package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/kirankothule/bookstore-users-api/domain/users"
	"github.com/kirankothule/bookstore-users-api/services"
	"github.com/kirankothule/bookstore-users-api/utils/errors"
)

func GetUser(c *gin.Context) {

	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if err != nil {
		restErr := errors.NewBadRequestError("user_id should be of type integer")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, result)

}

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err.Error())
		restErr := errors.NewBadRequestError(err.Error())
		c.JSON(restErr.Status, restErr)
		return
	}

	result, err := services.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func UpdateUser(c *gin.Context) {
	var user users.User

	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if err != nil {
		restErr := errors.NewBadRequestError("user_id should be of type integer")
		c.JSON(restErr.Status, restErr)
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err.Error())
		restErr := errors.NewBadRequestError(err.Error())
		c.JSON(restErr.Status, restErr)
		return
	}
	user.ID = userID

	isPartial := c.Request.Method == http.MethodPatch

	result, updateErr := services.UpdateUser(isPartial, user)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}
	c.JSON(http.StatusOK, result)
}
