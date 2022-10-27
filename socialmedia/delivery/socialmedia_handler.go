package delivery

import (
	authentication "mygram/authentication"
	"mygram/domain"
	authorize "mygram/socialmedia/delivery/authorization"
	"mygram/socialmedia/utility"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type socialMediaHandler struct {
	socialMediaUseCase domain.SocialMediaUseCase
}

func NewSocialMediaHandler(routers *gin.Engine, socialMediaUseCase domain.SocialMediaUseCase) {
	handler := &socialMediaHandler{socialMediaUseCase}
	router := routers.Group("/socialmedias")
	{
		router.Use(authentication.Authentication())
		router.GET("", handler.Fetch)
		router.POST("", handler.Store)
		router.PUT("/:socialMediaId", authorize.Authorization(handler.socialMediaUseCase), handler.Update)
		router.DELETE("/:socialMediaId", authorize.Authorization(handler.socialMediaUseCase), handler.Delete)
	}
}

func (handler *socialMediaHandler) Fetch(ctx *gin.Context) {
	var socialMedias []domain.SocialMedia
	var err error
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))
	err = handler.socialMediaUseCase.Fetch(ctx.Request.Context(), &socialMedias, userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"socialMedia": socialMedias,
	})
}
func (handler *socialMediaHandler) Store(ctx *gin.Context) {
	var socialMedia domain.SocialMedia
	var err error
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))
	if err = ctx.ShouldBindJSON(&socialMedia); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}
	socialMedia.UserID = userID
	err = handler.socialMediaUseCase.Store(ctx.Request.Context(), &socialMedia)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, utility.NewSocialMedia{
		ID:             socialMedia.ID,
		UserID:         socialMedia.UserID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		CreatedAt:      socialMedia.CreatedAt,
	})
}

func (handler *socialMediaHandler) Update(ctx *gin.Context) {
	var (
		socialMedia domain.SocialMedia
		err         error
	)
	socialMediaID := ctx.Param("socialMediaId")
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))
	if err = ctx.ShouldBindJSON(&socialMedia); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	updatedSocialMedia := domain.SocialMedia{
		UserID:         userID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
	}
	socialMedia, err = handler.socialMediaUseCase.Update(ctx.Request.Context(), updatedSocialMedia, socialMediaID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, utility.UpdatedSocialMedia{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserID:         socialMedia.UserID,
		UpdatedAt:      socialMedia.UpdatedAt,
	})
}

func (handler *socialMediaHandler) Delete(ctx *gin.Context) {
	socialMediaID := ctx.Param("socialMediaId")

	if err := handler.socialMediaUseCase.Delete(ctx.Request.Context(), socialMediaID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "your social media has been successfully deleted",
	})
}
