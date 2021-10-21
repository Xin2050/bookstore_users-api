package usersController

import (
	"bookstore_users-api/src/domain/users"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
func CreateUser(c *gin.Context) {
	var user users.User
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error(), "Other": 1234})
		return
	}
	c.JSON(200, &user)
}
