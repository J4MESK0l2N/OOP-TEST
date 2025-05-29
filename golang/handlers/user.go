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

	for _, user := range user_list {
		if(user.Username == u.Username) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"code":    500,
				"message": "Username duplicate.",
			})
			return
		}
	}

	if u.Username == "" || u.Password == ""  {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    400,
			"message": "Require Username and Password",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    500,
			"message": "Password hashing failed",
		})
		return
	}

	u.Password = string(hashedPassword)
	u.ID = len(user_list) + 1 

	user_list = append(user_list, u)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       u.ID,
		"username": u.Username,
		"role":     u.Role,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    500,
			"message": "Failed to generate token",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"user":  u.PublicUser(),
		"token": tokenString,
	})
}