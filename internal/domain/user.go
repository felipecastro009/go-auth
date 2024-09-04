package domain

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       string `gorm:"type:char(36);primaryKey"`
	Name     string `gorm:"size:100;not null"`
	Email    string `gorm:"size:100;unique;not null"`
	Password string `gorm:"not null"`
}

func NewUser(name, email, password string) *User {
	return &User{
		ID:       uuid.NewString(),
		Name:     name,
		Email:    email,
		Password: password,
	}
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.NewString()
	}
	if u.Password != "" {
		hash, err := MakePassword(u.Password)
		if err != nil {
			return nil
		}
		u.Password = hash
	}
	return
}

func MakePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
