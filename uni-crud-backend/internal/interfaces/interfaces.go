package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/worldwidepaniel/uni-crud-project/internal/config"
)

type Config interface {
	InitializeConfig() config.Config
}

type JWTUtils interface {
	GenerateJWT(user string, secret []byte) (string, error)
	ValidateJWT(tokenString string, secret []byte) error
}

type Router interface {
	InitializeRouter(*config.Config, JWTUtils) *gin.Engine
}
