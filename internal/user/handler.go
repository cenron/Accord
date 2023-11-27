package user

/*
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

func (u *UserHandler) Signup(c *gin.Context) {
	var req SignupUserReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := u.service.Signup(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (u *UserHandler) Login(c *gin.Context) {
	var req LoginUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	res, err := u.service.Login(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.SetCookie("jwt", res.accessToken, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, res)
}

func (u *UserHandler) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logged out user."})
}
*/
