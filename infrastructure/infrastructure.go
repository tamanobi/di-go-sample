package infrastructure

import "log"

type MySQL string

type HogeInfrastructure struct {
	db MySQL
}

func NewHogeInfrastructure() *HogeInfrastructure {
	var t MySQL
	t = "MySQL"
	log.Printf("connect DB(%s) on HogeInfrastructure\n", t)
	return &HogeInfrastructure{
		t,
	}
}

func (_ HogeInfrastructure) AddHoge() {
	log.Println("AddHoge on HogeInfrastructure")
}
