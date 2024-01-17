package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type JSON json.RawMessage

// Scan value into Jsonb, implements sql.Scanner interface
func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}

func (JSON) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "mysql", "sqlite":
		return "JSON"
	case "postgres":
		return "JSONB"
	}
	return ""
}

type Room struct {
	ID              int            `gorm:"column:id;primaryKey;autoIncrement"`
	RoomName        string         `gorm:"column:roomName;unique;not null;index"`
	RoomType        string         `gorm:"column:roomType;not null;index"` //single, double, suite, deluxe, etc
	Capacity        JSON           `gorm:"column:capacity;not null"`       // format { adults: 2, children: 1 }
	IsAvailable     bool           `gorm:"column:isAvailable;default:TRUE"`
	OccupancyStatus string         `gorm:"column:occupancyStatus;enum('vacant', 'occupied', 'undergoing cleaning/maintenance');default:'vacant';not null"`
	Amenities       JSON           `gorm:"column:amenities;"`
	View            string         `gorm:"column:view"`
	Price           JSON           `gorm:"column:price;not null"` //format {amount:50, currency:"US dollar" }
	Policy          string         `gorm:"column:policy"`
	AdditionalNotes string         `gorm:"column:additionalNotes"`
	Publish         JSON           `gorm:"column:publish;default:'{\"isPublished\": false, \"publishedAt\": \"\"}'"`
	RoomImages      []RoomImage    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt       time.Time      `gorm:"column:createdAt"`
	UpdatedAt       time.Time      `gorm:"column:updatedAt"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deletedAt;index"`
}

type RoomImage struct {
	ID        int            `gorm:"column:id;primaryKey;autoIncrement"`
	RoomID    int            `gorm:"column:roomId;not null; index"`
	ViewType  string         `gorm:"column:viewType;enum('interior', 'exterior', 'bathroom');default:'interior';not null"`
	URL       string         `gorm:"column:url;not null"`
	Path      string         `gorm:"column:path;not null"`
	CreatedAt time.Time      `gorm:"column:createdAt"`
	UpdatedAt time.Time      `gorm:"column:updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deletedAt;index"`
}
