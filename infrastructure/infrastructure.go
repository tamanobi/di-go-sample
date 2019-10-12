package infrastructure

import "log"

type MySQL string

type HogeInfrastructure struct {
	db MySQL
}

func NewHogeInfrastructure(db *MySQL) *HogeInfrastructure {
	log.Printf("connect DB(%s) on HogeInfrastructure\n", *db)
	return &HogeInfrastructure{
		*db,
	}
}

func (_ HogeInfrastructure) AddHoge() {
	log.Println("AddHoge on HogeInfrastructure")
}
