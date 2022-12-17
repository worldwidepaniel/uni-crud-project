package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/worldwidepaniel/uni-crud-project/internal/interfaces"
)

func JWTAuth(secret []byte, jwtUtils interfaces.JWTUtils) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtToken := ctx.Request.Header.Get("token")
		if !strings.Contains(jwtToken, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			ctx.Abort()
		}
		jwtToken = strings.TrimPrefix(jwtToken, "Bearer ")
		if err := jwtUtils.ValidateJWT(jwtToken, secret); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
		}
		ctx.Next()
	}
}
