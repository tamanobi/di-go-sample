package service

import (
	"log"

	"github.com/tamanobi/di-go-sample/infrastructure"
)

type HogeService struct {
	hi infrastructure.HogeInfrastructure
}

func NewHogeService(hi *infrastructure.HogeInfrastructure) *HogeService {
	return &HogeService{*hi}
}

func (hs HogeService) AddHoge() {
	log.Println("AddHoge on HogeService")
	hs.hi.AddHoge()
}
