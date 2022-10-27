package middleware

import (
	"fmt"
	"mygram/domain"
	"mygram/helpers"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authorization(photoUseCase domain.PhotoUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var photo domain.Photo
		var err error
		photoID := ctx.Param("photoId")
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := string(userData["id"].(string))
		err = photoUseCase.GetByID(ctx.Request.Context(), &photo, photoID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, helpers.ResponseMessage{
				Status:  "fail",
				Message: fmt.Sprintf("photo with id %s doesn't exist", photoID),
			})
			return
		}
		if photo.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helpers.ResponseMessage{
				Status:  "unauthorized",
				Message: "you don't have permission to view or edit this photo",
			})
			return
		}
	}
}
