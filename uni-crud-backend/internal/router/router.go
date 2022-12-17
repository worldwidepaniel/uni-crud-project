package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/worldwidepaniel/uni-crud-project/internal/interfaces"
	"github.com/worldwidepaniel/uni-crud-project/internal/middleware"
	"github.com/worldwidepaniel/uni-crud-project/internal/structs"
)

type Router struct{}

func (r Router) InitializeRouter(cfg *structs.Config, jwtUtils interfaces.JWTUtils) *gin.Engine {
	router := gin.Default()

	router.POST("/login", func(c *gin.Context) {
		var payload structs.Login
		if err := c.BindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Malformed request",
			})
			return
		}
		jwtToken, err := jwtUtils.GenerateJWT(payload.Username, []byte(cfg.JWTSecret))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "no user found with given credentials",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": jwtToken,
		})
	})

	authorized := router.Group("/api")
	authorized.Use(middleware.JWTAuth([]byte(cfg.JWTSecret), jwtUtils))
	{
		authorized.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"ping": "pong",
			})
		})
	}

	return router
}
