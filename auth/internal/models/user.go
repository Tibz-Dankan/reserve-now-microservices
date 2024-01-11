// package user
package models

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Tibz-Dankan/reserve-now-microservices/internal/config"
	"golang.org/x/crypto/bcrypt"
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
	ID                     int            `gorm:"column:id;primaryKey;autoIncrement"`
	Name                   string         `gorm:"column:name"`
	Email                  string         `gorm:"column:email;uniqueIndex:compositeindex"`
	Password               string         `gorm:"column:password"`
	Country                string         `gorm:"column:country"`
	PasswordResetToken     *string        `gorm:"column:passwordResetToken;index"`
	PasswordResetExpiresAt *time.Time     `gorm:"column:passwordResetExpiresAt"`
	CreatedAt              time.Time      `gorm:"column:createdAt"`
	UpdatedAt              time.Time      `gorm:"column:updatedAt"`
	DeletedAt              gorm.DeletedAt `gorm:"index"`
}

var db = config.Db()

func DBAutoMigrate() {
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Failed to make auto migration", err)
	}
	fmt.Println("Auto Migration successful")
}

// Hash password before creating user
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return
}

func (u *User) Create(user User) (int, error) {
	result := db.Create(&user)

	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}

func (u *User) FindOne(id int) (User, error) {
	var user User
	db.First(&user, id)

	return user, nil
}

func (u *User) FindByEMail(email string) (User, error) {
	var user User
	db.First(&user, "email = ?", email)

	return user, nil
}

func (u *User) FindByPasswordResetToken(passwordResetToken string) (User, error) {
	var user User
	db.First(&user, "passwordResetToken = ?", passwordResetToken)

	return user, nil
}

func (u *User) FindAll() ([]User, error) {
	var users []User
	db.Find(&users)

	return users, nil
}

// Update updates one user in the database, using the information
// stored in the receiver u
func (u *User) Update() error {
	db.Save(&u)

	return nil
}

func (u *User) Delete(id int) error {
	db.Delete(&User{}, id)

	return nil
}

// ResetPassword is the method we will use to change a user's password.
func (u *User) ResetPassword(password string) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	db.Model(&User{}).Where("id = ?", u.ID).Update("password", hashedPassword)

	return nil
}

func (u *User) PasswordMatches(plainTextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainTextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}