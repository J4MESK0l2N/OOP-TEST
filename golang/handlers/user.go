package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"test-oop-golang/models"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var user_list []models.User

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var list []map[string]interface{}
	
	for _, u := range user_list {
		list = append(list, u.PublicUser())
	}

	json.NewEncoder(w).Encode(list)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	
	var u models.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	
	if u.Role == "" {
		u.Role = "user"
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		http.Error(w, "Password hashing failed", http.StatusInternalServerError)
		return
	}

	u.Password = string(hashedPassword)
	u.ID = len(users) + 1

	users = append(users, u)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       u.ID,
		"username": u.Username,
		"role":     u.Role,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"user":  u.PublicUser(),
		"token": tokenString,
	})
}