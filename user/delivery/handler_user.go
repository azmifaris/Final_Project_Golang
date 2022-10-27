package delivery

import (
	authentication "mygram/authentication"
	"mygram/domain"
	"mygram/helpers"
	"mygram/user/utility"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type userHandler struct {
	userUseCase domain.UserUseCase
}

func NewUserHandler(routers *gin.Engine, userUseCase domain.UserUseCase) {
	handler := &userHandler{userUseCase}

	router := routers.Group("/users")
	{
		router.POST("/register", handler.Register)
		router.POST("/login", handler.Login)
		router.PUT("/", authentication.Authentication(), handler.Update)
		router.DELETE("/", authentication.Authentication(), handler.Delete)
	}
}

func (route *userHandler) Register(c *gin.Context) {
	var user domain.User
	var err error
	err = c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	err = route.userUseCase.Register(c.Request.Context(), &user)
	if err != nil {
		if strings.Contains(err.Error(), "id_users_username") {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"error":   "Conflict",
				"message": "the username already used",
			})
			return
		}
		if strings.Contains(err.Error(), "id_users_email") {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"error":   "Conflict",
				"message": "the email already used",
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":       user.ID,
		"email":    user.Email,
		"username": user.Username,
		"age":      user.Age,
	})
}

func (route *userHandler) Login(c *gin.Context) {
	var user domain.User
	var err error
	var token string
	err = c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	err = route.userUseCase.Login(c.Request.Context(), &user)
	if err != nil {
		if strings.Contains(err.Error(), "invalid credential") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": err.Error(),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	token, err = helpers.GenerateToken(user.ID, user.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (route *userHandler) Update(c *gin.Context) {
	var user domain.User
	var err error
	userData := c.MustGet("userData").(jwt.MapClaims)
	_ = string(userData["id"].(string))
	err = c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	updatedUser := domain.User{
		Username: user.Username,
		Email:    user.Email,
	}
	user, err = route.userUseCase.Update(c.Request.Context(), updatedUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, utility.UserUpdated{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		Age:       user.Age,
		UpdatedAt: user.UpdatedAt,
	})
}

func (route *userHandler) Delete(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))
	err := route.userUseCase.Delete(c, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	},
	)
}
