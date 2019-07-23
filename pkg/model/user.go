package model

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// User is the structure of a user
type User struct {
	Email              string   `json:"email" bson:"email"`
	EmailBackup        []string `json:"-" bson:"emailBackup"`
	ID                 string   `json:"id" bson:"_id"`
	Name               string   `json:"name" bson:"name"`
	Password           string   `json:"password" bson:"password"`
	ResetPasswordToken string   `json:"resetPasswordToken" bson:"resetPasswordToken"`
	Token              string   `json:"token" bson:"-"`
	VerifyToken        string   `json:"verifyToken" bson:"verifyToken"`
}

// HashPassword hashed the password of the user
func (u *User) HashPassword() error {
	key, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(key)

	return nil
}

// MatchPassword returns true if the hashed user password matches the password
func (u *User) MatchPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err == nil {
		return true
	}

	return false
}
