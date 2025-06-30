package controller

import (
	"backend2/models"
	"backend2/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Description Create new user
// @Tags createuser
// @Accept json
// @Produce json
// @Param user body models.User true "User Data"
// @Success 201 {object} models.ResponseUser
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /users [post]
func CreateUser(c *gin.Context) {
	user := models.User{}
	c.ShouldBind(&user)

	if user.Email == "" || user.Name == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "name, email and/or password is missing",
		})
		return
	}

	err := models.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Create user success!",
		Result: models.ResponseUser{
			Name:  user.Name,
			Email: user.Email,
		},
	})
}

// @Description List all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {string} string "string"
// @Failure 500 {object} utils.Response
// @Router /users [get]
func GetAllUsers(c *gin.Context) {
	users, err := models.FindAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to get data",
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Get data all users success!",
		PageInfo: map[string]any{
			"totalData": len(users),
		},
		Result: users,
	})
}

// @Description Update user
// @Tags update
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "User Data"
// @Success 201 {object} models.ResponseUser
// @Failure 500 {object} utils.Response
// @Router /users/{id} [patch]
func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.User{}
	c.ShouldBind(&user)

	err := models.UpdateUser(id, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to update data",
		})
		return
	}
	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Success to update data",
		Result: models.ResponseUser{
			Name:  user.Name,
			Email: user.Email,
		},
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := models.FindUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to get data",
		})
		return
	}
	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Success to get data",
		Result:  user,
	})
}

// @Description Update user
// @Tags delete
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to delete user",
		})
		return
	}
	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Delete user success",
	})
}
