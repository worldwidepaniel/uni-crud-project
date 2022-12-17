package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/worldwidepaniel/uni-crud-project/internal/structs"
)

type Config interface {
	InitializeConfig() structs.Config
}

type JWTUtils interface {
	GenerateJWT(user string, secret []byte) (string, error)
	ValidateJWT(tokenString string, secret []byte) error
}

type Router interface {
	InitializeRouter(*structs.Config, JWTUtils) *gin.Engine
}
