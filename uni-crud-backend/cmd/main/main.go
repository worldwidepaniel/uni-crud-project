package main

import (
	"fmt"

	"github.com/worldwidepaniel/uni-crud-project/internal/config"
	"github.com/worldwidepaniel/uni-crud-project/internal/router"
)

func main() {
	cfg, err := config.InitializeConfig()
	if err != nil {
		fmt.Println(err)
	}

	router := router.InitializeRouter(cfg)
	router.Run()
}
