package models

import (
	"github.com/Z0marlin/wocbackend/db"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User is a generic website user.
type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Password  string
}

// UserModel is a placeholder for all functions related to User model.
type UserModel struct{}

// NewUser creates a new user
func (um UserModel) NewUser(fname string, lname string, pass string) User {
	return User{
		FirstName: fname,
		LastName:  lname,
		Password:  pass,
	}
}

// ListUsers lists all users
func (um UserModel) ListUsers(limit int) ([]User, error) {
	var users []User
	res := db.D().Limit(limit).Find(&users)
	return users, res.Error
}

// GetUserByID returns user with given ID
func (um UserModel) GetUserByID(id string) (User, error) {
	var user User
	res := db.D().First(&user, id)
	return user, res.Error
}

// CreateUser creates a new user entry in the database
func (um UserModel) CreateUser(u *User) error {
	res := db.D().Create(u)
	return res.Error
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(u.Password),
		bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(pass)
	return nil
}
