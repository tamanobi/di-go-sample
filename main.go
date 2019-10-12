package main

import (
	"github.com/tamanobi/di-go-sample/infrastructure"
	"github.com/tamanobi/di-go-sample/service"
)

func main() {
	var mysql infrastructure.MySQL
	mysql = "MySQL"
	hi := infrastructure.NewHogeInfrastructure(&mysql)
	hs := service.NewHogeService(hi)
	hs.AddHoge()
}
