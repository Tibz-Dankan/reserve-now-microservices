package models

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Tibz-Dankan/reserve-now-microservices/room/internal/config"
)

var db = config.Db()

func DBAutoMigrate() {
	err := db.AutoMigrate(&Room{}, &RoomImage{})
	if err != nil {
		log.Fatal("Failed to make auto migration", err)
	}
	fmt.Println("Auto Migration successful")
}

func (r *Room) getRoomCapacityValue(key string) (int, bool) {

	var capacityMap map[string]int
	capacity := r.Capacity

	err := json.Unmarshal(capacity, &capacityMap)
	if err != nil {
		return 0, false
	}

	value, exists := capacityMap[key]
	return value, exists
}

func (r *Room) IsValidRoomCapacity() bool {
	adults, adultsExist := r.getRoomCapacityValue("adults")
	children, childrenExist := r.getRoomCapacityValue("children")

	if !adultsExist || adults <= 0 {
		return false
	}
	if childrenExist && children <= 0 {
		return false
	}

	return true
}

func (r *Room) getRoomPriceValue(key string) (int, bool) {

	var priceMap map[string]int
	price := r.Price

	err := json.Unmarshal(price, &priceMap)
	if err != nil {
		return 0, false
	}

	value, exists := priceMap[key]
	return value, exists
}

func (r *Room) IsValidRoomPrice() bool {
	amount, amountExist := r.getRoomPriceValue("amount")
	currency, currencyExist := r.getRoomPriceValue("currency")

	if !amountExist || !currencyExist || amount <= 0 || currency <= 0 {
		return false
	}

	return true
}

func (r *Room) Create(room Room) (int, error) {
	result := db.Create(&room)

	if result.Error != nil {
		return 0, result.Error
	}
	return room.ID, nil
}

func (r *Room) FindOne(id int) (Room, error) {
	var room Room
	err := db.First(&room, id).Preload("RoomImages").Error
	if err != nil {
		return room, err
	}

	return room, nil
}

func (r *Room) FindByName(name string) (Room, error) {
	var room Room
	err := db.First(&room, "roomName = ?", name).Preload("RoomImages").Error
	if err != nil {
		return room, err
	}

	return room, nil
}

func (r *Room) FindAll() ([]Room, error) {
	var rooms []Room
	err := db.Find(&rooms).Preload("RoomImages").Error
	if err != nil {
		return rooms, err
	}

	return rooms, nil
}

func (r *Room) Update(id int) error {
	r.ID = id
	db.Save(&r)

	return nil
}

func (r *Room) Delete(id int) error {
	// TODO: Delete room beds
	db.Delete(&RoomImage{}, "roomId = ?", r.ID)
	db.Delete(&Room{}, id)

	return nil
}

func (r *Room) IsPublished() bool {

	var publish map[string]bool
	err := json.Unmarshal(r.Publish, &publish)
	if err != nil {
		fmt.Println("error : ", err)
	}

	isPublished := publish["isPublished"]

	return isPublished
}

func (r *Room) UpdateAsPublished(id int) error {

	r.Publish = JSON([]byte(`{"isPublished": true, "isPublishedAt": time.now()}`))

	err := db.Model(&Room{}).Where("id = ?", id).Update("publish", r.Publish).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Room) UpdateAsUnPublished(id int) error {

	r.Publish = JSON([]byte(`{"isPublished": false, "isPublishedAt": ""}`))

	db.Model(&Room{}).Where("id = ?", id).Update("publish", r.Publish)

	return nil
}
