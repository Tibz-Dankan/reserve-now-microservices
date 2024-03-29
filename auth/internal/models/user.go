package models

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Tibz-Dankan/reserve-now-microservices/auth/internal/config"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Role string

const (
	Admin  Role = "admin"
	Client Role = "client"
	Staff  Role = "staff"
)

type User struct {
	ID                     int            `gorm:"column:id;primaryKey;autoIncrement"`
	Name                   string         `gorm:"column:name;not null"`
	Email                  string         `gorm:"column:email;unique;not null"`
	Password               string         `gorm:"column:password;not null"`
	Country                string         `gorm:"column:country;not null"`
	PasswordResetToken     string         `gorm:"column:passwordResetToken;index"`
	PasswordResetExpiresAt time.Time      `gorm:"column:passwordResetExpiresAt"`
	Role                   string         `gorm:"column:role;enum('admin', 'client', 'staff');default:'client';not null"`
	CreatedAt              time.Time      `gorm:"column:createdAt"`
	UpdatedAt              time.Time      `gorm:"column:updatedAt"`
	DeletedAt              gorm.DeletedAt `gorm:"column:deletedAt;index"`
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
	db.Model(&User{}).Where("id = ?", u.ID).Update("\"passwordResetExpiresAt\"", time.Now())

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

func (u *User) ValidRole(role string) bool {
	roles := []string{"admin", "client", "staff"}

	for _, r := range roles {
		if r == role {
			return true
		}
	}

	return false
}

func (u *User) SetRole(role string) error {
	isValidRole := u.ValidRole(role)

	if !isValidRole {
		return errors.New("invalid user role")
	}

	u.Role = role
	return nil
}

func (u *User) FindByPasswordResetToken(resetToken string) (User, error) {
	var user User
	hashedToken := sha256.New()
	hashedToken.Write([]byte(resetToken))
	hashedTokenByteSlice := hashedToken.Sum(nil)
	hashedTokenString := hex.EncodeToString(hashedTokenByteSlice)

	result := db.Where("\"passwordResetToken\" = ? AND \"passwordResetExpiresAt\" > ?", hashedTokenString, time.Now()).Find(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (u *User) CreatePasswordResetToken() (string, error) {
	resetToken := uuid.NewString()

	hashedToken := sha256.New()
	hashedToken.Write([]byte(resetToken))
	hashedTokenByteSlice := hashedToken.Sum(nil)
	hashedTokenString := hex.EncodeToString(hashedTokenByteSlice)

	u.PasswordResetToken = hashedTokenString
	u.PasswordResetExpiresAt = time.Now().Add(20 * time.Minute)

	db.Save(&u)

	return resetToken, nil
}
