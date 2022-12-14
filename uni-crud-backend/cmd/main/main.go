package main

import (
	"github.com/worldwidepaniel/uni-crud-project/internal/config"
	"github.com/worldwidepaniel/uni-crud-project/internal/interfaces"
	"github.com/worldwidepaniel/uni-crud-project/internal/jwtUtils"
	"github.com/worldwidepaniel/uni-crud-project/internal/router"
	"github.com/worldwidepaniel/uni-crud-project/internal/structs"
)

type service struct {
	// config   interfaces.Config
	jwtutils interfaces.JWTUtils
	router   interfaces.Router
}

func NewService(jwt *jwtUtils.JWTUtils, router *router.Router) *service {
	return &service{
		// config:   cfg,
		jwtutils: jwt,
		router:   router,
	}
}

func (srv *service) Run(cfg *structs.Config) {
	r := srv.router.InitializeRouter(cfg, srv.jwtutils)
	r.Run()
}

func main() {
	cfg := config.InitializeConfig()
	jwt := jwtUtils.JWTUtils{}
	r := router.Router{}
	service := NewService(&jwt, &r)
	service.Run(&cfg)
}
