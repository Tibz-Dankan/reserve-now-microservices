package models

func (r *RoomImage) Create(roomImage RoomImage) (int, error) {
	err := db.Create(&roomImage).Error

	if err != nil {
		return 0, err
	}

	return roomImage.ID, nil
}

func (ri *RoomImage) FindOne(id int) (*RoomImage, error) {
	var roomImage *RoomImage
	err := db.First(&roomImage, id).Error
	if err != nil {
		return roomImage, err
	}

	return roomImage, nil
}

func (ri *RoomImage) FindAll() ([]RoomImage, error) {
	var roomImages []RoomImage
	err := db.Find(&roomImages).Error
	if err != nil {
		return roomImages, err
	}

	return roomImages, nil
}

func (ri *RoomImage) Update(id int) error {
	ri.ID = id
	if err := db.Save(&ri).Error; err != nil {
		return err
	}

	return nil
}

func (ri *RoomImage) Delete(id int) error {
	err := db.Delete(&RoomImage{}, "id = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}

func (ri *RoomImage) DeleteByRoomId(roomId int) error {
	err := db.Delete(&RoomImage{}, "roomId = ?", roomId).Error
	if err != nil {
		return err
	}

	return nil
}
