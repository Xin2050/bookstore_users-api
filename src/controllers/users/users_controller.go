package usersController

import (
	"bookstore_users-api/src/domain/users"
	"bookstore_users-api/src/services"
	"bookstore_users-api/src/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, userErr:= strconv.ParseInt(c.Param("user_id"),10,64)
	if userErr != nil {
		err:= errors.BadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr !=nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
func CreateUser(c *gin.Context) {
	var user users.User
	if err :=c.ShouldBindJSON(&user);err != nil {
		restError := errors.BadRequestError(err.Error())
		c.JSON(restError.Status, restError)
		return
	}
	newUser, err := services.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, newUser)
}
