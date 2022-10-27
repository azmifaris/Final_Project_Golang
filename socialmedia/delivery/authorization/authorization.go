package middleware

import (
	"fmt"
	"mygram/domain"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authorization(socialMediaUseCase domain.SocialMediaUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var socialMedia domain.SocialMedia
		var err error
		socialMediaID := ctx.Param("socialMediaId")
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := string(userData["id"].(string))
		if err = socialMediaUseCase.GetByID(ctx.Request.Context(), &socialMedia, socialMediaID); err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": fmt.Sprintf("social media with id %s doesn't exist", socialMediaID),
			})
			return
		}
		if socialMedia.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you don't have permission to view or edit this social media",
			})
			return
		}
	}
}
