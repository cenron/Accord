package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	service *UserService
}

func NewUserHandler(store UserStore) *UserHandler {
	return &UserHandler{
		service: NewUserService(store),
	}
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	var req CreateUserReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := u.service.CreateUser(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
