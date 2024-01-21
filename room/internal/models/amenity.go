package models

func (a *Amenity) Create(amenity Amenity) (int, error) {
	err := db.Create(&amenity).Error

	if err != nil {
		return 0, err
	}

	return amenity.ID, nil
}

func (a *Amenity) FindOne(id int) (*Amenity, error) {
	var amenity *Amenity
	err := db.First(&amenity, id).Error
	if err != nil {
		return amenity, err
	}

	return amenity, nil
}

func (a *Amenity) FindByItem(item string) (*Amenity, error) {
	var amenity *Amenity
	err := db.First(&amenity, "item = ?", item).Error
	if err != nil {
		return amenity, err
	}

	return amenity, nil
}

func (a *Amenity) FindAll() ([]Amenity, error) {
	var amenities []Amenity
	err := db.Find(&amenities).Error
	if err != nil {
		return amenities, err
	}

	return amenities, nil
}

func (a *Amenity) Update(id int) error {
	a.ID = id
	if err := db.Save(&a).Error; err != nil {
		return err
	}

	return nil
}

func (a *Amenity) Delete(id int) error {
	err := db.Delete(&Amenity{}, "id = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}
