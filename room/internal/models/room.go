package models

import (
	"fmt"
	"log"

	"github.com/Tibz-Dankan/reserve-now-microservices/room/internal/config"
)

// type Room struct {
// 	gorm.Model
// 	ID int `gorm:"column:id;primaryKey;autoIncrement"`
// }

var db = config.Db()

func DBAutoMigrate() {
	err := db.AutoMigrate(&Room{}, &RoomImage{})
	if err != nil {
		log.Fatal("Failed to make auto migration", err)
	}
	fmt.Println("Auto Migration successful")
}
