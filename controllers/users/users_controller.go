package users

import (
	"github.com/eremitic/bookstore_users-api/domain/users"
	"github.com/eremitic/bookstore_users-api/services"
	"github.com/eremitic/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"strconv"

	"net/http"
)

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("id"), 10, 64)

	if userErr != nil {
		err := errors.NewBadReqErr("id is number only")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)

	if getErr != nil {
		err := errors.NewNotFoundErr("user not found")
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, user)

}

func CreateUser(c *gin.Context) {

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {

		restErr := errors.NewBadReqErr("invalid json body")
		c.JSON(restErr.Status, restErr)

		return
	}

	result, saveErr := services.CreateUser(user)

	if saveErr != nil {

		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)

}

func UpdateUser(c *gin.Context) {

	var user users.User

	userId, userErr := strconv.ParseInt(c.Param("id"), 10, 64)

	if userErr != nil {
		err := errors.NewBadReqErr("id is number only")
		c.JSON(err.Status, err)
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {

		restErr := errors.NewBadReqErr("invalid json body")
		c.JSON(restErr.Status, restErr)

		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, saveErr := services.UpdateUser(isPartial, user)

	if saveErr != nil {

		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusOK, result)

}

func SearchUser(c *gin.Context) {

}
