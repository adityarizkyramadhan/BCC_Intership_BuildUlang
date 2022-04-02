package user

import (
	"BCC_Intership_BuildUlang/domain"
	data "BCC_Intership_BuildUlang/handler"
	"BCC_Intership_BuildUlang/helper"
	repouser "BCC_Intership_BuildUlang/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type newUser struct {
	Name     string `json:"name" binding:"required"`
	Contact  string `json:"contact" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Address  string `json:"address" binding:"required"`
}

var database = repouser.NewUserRepository(data.DB)

func RegisterUser(c *gin.Context) {
	var userRegister newUser
	if err := c.BindJSON(&userRegister); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, helper.HelperMessage(false, err.Error(), nil))
	}
	saveUser := domain.User{
		Name:     userRegister.Name,
		Contact:  userRegister.Contact,
		Username: userRegister.Username,
		Password: userRegister.Password,
		Address:  userRegister.Address,
	}
	data, err := database.Create(&saveUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.HelperMessage(false, err.Error(), nil))
	}
	c.JSON(http.StatusCreated, helper.HelperMessage(true, "Register Success", data))
}

func DeleteUser(c *gin.Context) {
	id := c.MustGet("id").(string)
	if err := database.Delete(id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.HelperMessage(false, err.Error(), nil))
	}
	c.JSON(http.StatusAccepted, helper.HelperMessage(true, "Delete Success", nil))
}

func UpdateUser(c *gin.Context) {
	id := c.MustGet("id").(string)
	var userUpdate newUser
	if err := c.BindJSON(&userUpdate); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, helper.HelperMessage(false, err.Error(), nil))
	}
	updateUser := domain.User{
		Name:     userUpdate.Name,
		Contact:  userUpdate.Contact,
		Username: userUpdate.Username,
		Password: userUpdate.Password,
		Address:  userUpdate.Address,
	}
	data, err := database.Update(id, &updateUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, helper.HelperMessage(false, err.Error(), nil))
	}
	c.JSON(http.StatusAccepted, helper.HelperMessage(true, "Update Success", data))
}
