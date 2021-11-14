package usersController

import (
	"bookstore_users-api/src/domain/users"
	"bookstore_users-api/src/services"
	"bookstore_users-api/src/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func TestServiceInterface() {
	//services.UsersService
}

func GetUserId(userIdParam string) (int64, *errors.RestError) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return -1, errors.BadRequestError("user id should be a number")
	}
	return userId, nil
}
func GetUser(c *gin.Context) {
	userId, userErr := GetUserId(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}
	user, getErr := services.UsersService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("x-public") == "true"))
}
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.BadRequestError(err.Error())
		c.JSON(restError.Status, restError)
		return
	}
	newUser, err := services.UsersService.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, newUser)
}

func UpdateUser(c *gin.Context) {
	userId, userErr := GetUserId(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.BadRequestError(err.Error())
		c.JSON(restError.Status, restError)
		return
	}
	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch
	result, err := services.UsersService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("x-public") == "true"))
}
func DeleteUser(c *gin.Context) {
	userId, userErr := GetUserId(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}
	user := users.User{Id: userId}
	err := services.UsersService.DeleteUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
func FindUsersByStatus(c *gin.Context) {
	status := c.Query("status")

	users, err := services.UsersService.FindUsersByStatus(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("x-public") == "true"))
}
