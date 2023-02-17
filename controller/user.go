package controller

import (
	"net/http"

	"github.com/clean-architecture/entity"
	"github.com/clean-architecture/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService usecase.UserService
}

func NewUserController(userService usecase.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var input entity.User
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.userService.CreateUser(input.Name, input.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}
