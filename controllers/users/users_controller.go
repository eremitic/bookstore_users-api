package users

import (
	"github.com/eremitic/bookstore_users-api/domain/users"
	"github.com/eremitic/bookstore_users-api/services"
	"github.com/eremitic/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"strconv"

	"net/http"
)

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)

	if userErr != nil {
		err := errors.NewBadReqErr("id is number only")

		return 0, err
	}
	return userId, nil
}
func Get(c *gin.Context) {
	userId, idErr := getUserId(c.Param("id"))

	if idErr != nil {
		c.JSON(idErr.Status, idErr)
	}

	user, getErr := services.GetUser(userId)

	if getErr != nil {
		err := errors.NewNotFoundErr("user not found")
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, user)

}

func Create(c *gin.Context) {

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

func Update(c *gin.Context) {

	var user users.User

	userId, idErr := getUserId(c.Param("id"))

	if idErr != nil {
		c.JSON(idErr.Status, idErr)
	}

	if err := c.ShouldBindJSON(&user); err != nil {

		restErr := errors.NewBadReqErr("invalid json body")
		c.JSON(restErr.Status, err)

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

func Delete(c *gin.Context) {

	userId, idErr := getUserId(c.Param("id"))

	if idErr != nil {
		c.JSON(idErr.Status, idErr)
	}

	if err := services.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)

		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
