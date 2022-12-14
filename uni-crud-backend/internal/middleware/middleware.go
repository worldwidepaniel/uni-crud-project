package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/worldwidepaniel/uni-crud-project/internal/jwtUtils"
)

func ValidateJWT(secret []byte) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtToken := ctx.Request.Header.Get("token")
		if err := jwtUtils.ValidateJWT(jwtToken, secret); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
		}
		ctx.Next()
	}
}
