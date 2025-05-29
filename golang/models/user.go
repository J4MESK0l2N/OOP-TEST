package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID 			int
	Username 	string
	Password	string
	Role		string
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) PublicUser() map[string]interface{} {
	return map[string]interface{} {
		"id":		u.ID,
		"username": u.Username,
		"role":		u.Role,
	}
}