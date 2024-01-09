package user

import (
	"time"

	"github.com/Tibz-Dankan/reserve-now-microservices/internal/config"
	"gorm.io/gorm"
)

// type User struct {
// 	UserId                 int       `json:"id"`
// 	Name                   string    `json:"name,omitempty"`
// 	Email                  string    `json:"email"`
// 	Password               string    `json:"-"`
// 	Country                string    `json:"country"`
// 	PasswordResetToken     time.Time `json:"-"`
// 	PasswordResetExpiresAt time.Time `json:"-"`
// 	CreatedAt              time.Time `json:"createdAt"`
// 	UpdatedAt              time.Time `json:"updatedAt"`
// }

type User struct {
	gorm.Model
	ID                     uint           `gorm:"column:id;primaryKey;autoIncrement"`
	Name                   string         `gorm:"column:name"`
	Email                  string         `gorm:"column:email;index"`
	Password               string         `gorm:"column:password"`
	Country                string         `gorm:"column:country"`
	PasswordResetToken     *string        `gorm:"column:passwordResetToken;index"`
	PasswordResetExpiresAt *time.Time     `gorm:"column:passwordResetExpiresAt"`
	CreatedAt              time.Time      `gorm:"column:createdAt"`
	UpdatedAt              time.Time      `gorm:"column:updatedAt"`
	DeletedAt              gorm.DeletedAt `gorm:"index"`
}

var db = config.Db()

//  //Before create hook, hash the user password
// func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
// 	u.UUID = uuid.New()

// 	if u.Role == "admin" {
// 		return errors.New("invalid role")
// 	}
// 	return
// }

//  //Before save hook, hash the user password
// func (u *User) BeforeSave(tx *gorm.DB) (err error) {
// 	u.UUID = uuid.New()

// 	if u.Role == "admin" {
// 		return errors.New("invalid role")
// 	}
// 	return
// }

func (u *User) FindOne(userId int) (User, error) {
	var user User
	// db.First(&user, 10)
	db.First(&user, "id = ?", userId)

	return user, nil
}

func (u *User) FindAll() ([]User, error) {
	var user []User
	// result := db.Find(&user)
	db.Find(&user)

	return user, nil
}
