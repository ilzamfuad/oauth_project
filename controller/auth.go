package controller

import (
	"net/http"
	"nexmedis_project/model"
	"nexmedis_project/service"
	"nexmedis_project/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userService service.UserService
}

func NewAuthController(userService service.UserService) *AuthController {
	return &AuthController{
		userService: userService,
	}
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (ac *AuthController) Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := model.User{
		Username: input.Username,
		Password: input.Password,
	}

	err := ac.userService.CreateUser(u)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (ac *AuthController) Login(c *gin.Context) {

	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ac.userService.LoginCheck(input.Username, input.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

func (ac *AuthController) CurrentUser(c *gin.Context) {

	cu, exists := c.Get(string(utils.UserCtxKey))
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	userID := cu.(int)

	u, err := ac.userService.GetUserByID(userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}
