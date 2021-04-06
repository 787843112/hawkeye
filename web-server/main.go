package main

import (
	"web-server/config"
	"web-server/controller"
	"web-server/service"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config := config.InitConfig()
	s := service.New(config)
	controller.Init(s)
	defer s.Close()
}
