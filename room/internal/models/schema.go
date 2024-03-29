package models

import (
	"time"

	"gorm.io/gorm"
)

type Room struct {
	ID              int            `gorm:"column:id;primaryKey;autoIncrement"`
	RoomName        string         `gorm:"column:roomName;unique;not null;index"`
	RoomType        string         `gorm:"column:roomType;not null;index"` //single, double, suite, deluxe, etc
	IsAvailable     bool           `gorm:"column:isAvailable;default:TRUE"`
	OccupancyStatus string         `gorm:"column:occupancyStatus;enum('vacant', 'occupied', 'undergoing cleaning/maintenance');default:'vacant';not null"`
	View            string         `gorm:"column:view"`
	Policy          string         `gorm:"column:policy"`
	AdditionalNotes string         `gorm:"column:additionalNotes"`
	UpdatedByUserId int            `gorm:"column:updatedByUserId;not null"`
	Images          []RoomImage    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Capacity        RoomCapacity   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Price           RoomPrice      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Publicity       RoomPublicity  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Amenities       []RoomAmenity  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RoomBeds        []RoomBed      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt       time.Time      `gorm:"column:createdAt"`
	UpdatedAt       time.Time      `gorm:"column:updatedAt"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deletedAt;index"`
}

type RoomImage struct {
	ID              int            `gorm:"column:id;primaryKey;autoIncrement"`
	RoomID          int            `gorm:"column:roomId;not null;index"`
	ViewType        string         `gorm:"column:viewType;enum('interior', 'exterior', 'bathroom');default:'interior';not null"`
	URL             string         `gorm:"column:url;not null"`
	Path            string         `gorm:"column:path;not null"`
	UpdatedByUserId int            `gorm:"column:updatedByUserId;not null"`
	CreatedAt       time.Time      `gorm:"column:createdAt"`
	UpdatedAt       time.Time      `gorm:"column:updatedAt"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deletedAt;index"`
}

type RoomCapacity struct {
	ID              int            `gorm:"column:id;primaryKey;autoIncrement"`
	RoomID          int            `gorm:"column:roomId;not null;index"`
	Adults          int            `gorm:"column:adults;not null;"`
	Children        int            `gorm:"column:children;not null;"`
	UpdatedByUserId int            `gorm:"column:updatedByUserId;not null"`
	CreatedAt       time.Time      `gorm:"column:createdAt"`
	UpdatedAt       time.Time      `gorm:"column:updatedAt"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deletedAt;index"`
}

type RoomPrice struct {
	ID              int            `gorm:"column:id;primaryKey;autoIncrement"`
	RoomID          int            `gorm:"column:roomId;not null;index"`
	Amount          int            `gorm:"column:amount;not null;"`
	Currency        string         `gorm:"column:currency;not null;"`
	UpdatedByUserId int            `gorm:"column:updatedByUserId;not null"`
	CreatedAt       time.Time      `gorm:"column:createdAt"`
	UpdatedAt       time.Time      `gorm:"column:updatedAt"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deletedAt;index"`
}

type RoomPublicity struct {
	ID              int            `gorm:"column:id;primaryKey;autoIncrement"`
	RoomID          int            `gorm:"column:roomId;not null;index"`
	IsPublished     bool           `gorm:"column:isPublished;default:FALSE;"`
	UpdatedByUserId int            `gorm:"column:updatedByUserId;not null"`
	CreatedAt       time.Time      `gorm:"column:createdAt"`
	UpdatedAt       time.Time      `gorm:"column:updatedAt"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deletedAt;index"`
}

type Amenity struct {
	ID              int            `gorm:"column:id;primaryKey;autoIncrement"`
	Item            string         `gorm:"column:item;unique;not null;"`
	RoomAmenities   []RoomAmenity  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UpdatedByUserId int            `gorm:"column:updatedByUserId;not null"`
	CreatedAt       time.Time      `gorm:"column:createdAt"`
	UpdatedAt       time.Time      `gorm:"column:updatedAt"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deletedAt;index"`
}
type RoomAmenity struct {
	ID              int            `gorm:"column:id;primaryKey;autoIncrement"`
	RoomID          int            `gorm:"column:roomId;not null;index"`
	AmenityID       int            `gorm:"column:amenityId;not null;index"`
	UpdatedByUserId int            `gorm:"column:updatedByUserId;not null"`
	CreatedAt       time.Time      `gorm:"column:createdAt"`
	UpdatedAt       time.Time      `gorm:"column:updatedAt"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deletedAt;index"`
}

type RoomBed struct {
	ID              int            `gorm:"column:id;primaryKey;autoIncrement"`
	RoomID          int            `gorm:"column:roomId;not null;index"`
	BedType         string         `gorm:"column:bedType;not null;"`
	UpdatedByUserId int            `gorm:"column:updatedByUserId;not null"`
	CreatedAt       time.Time      `gorm:"column:createdAt"`
	UpdatedAt       time.Time      `gorm:"column:updatedAt"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deletedAt;index"`
}
