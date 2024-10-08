package models

import (
	"golang.org/x/crypto/bcrypt"
)

type UserRole string

const (
	Admin   UserRole = "Admin"
	Student UserRole = "Student"
	Teacher UserRole = "Teacher"
)

type User struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Role     UserRole `json:"role"`
}

func (u User) GetID() string {
	return u.Username
}

// HashPassword hashes the password for storing
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword checks the provided password against the stored hashed password
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
