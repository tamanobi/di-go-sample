package main

import (
	"github.com/tamanobi/di-go-sample/service"
)

func main() {
	hs := service.NewHogeService()
	hs.AddHoge()
}
