package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"test-oop-golang/models"

	"github.com/golang-jwt/jwt/v5"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	for _, u := range user_list {

		fmt.Println(u.Username)
		fmt.Println(input.Username)
		fmt.Println(u.Username)
		fmt.Println(input.Password)
		fmt.Println(u.CheckPassword(input.Password))

		if(u.Username == input.Username && u.CheckPassword(input.Password)) {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"id":		u.ID,
				"username":	u.Username,
				"role":		u.Role,
				"exp":		time.Now().Add(time.Hour * 1).Unix(),
			})

			secret := os.Getenv("JWT_SECRET")
			tokenString, err := token.SignedString([]byte(secret))

			if err != nil {
				http.Error(w, "Generate Token failed", http.StatusInternalServerError)
				return
			}

			fmt.Println(tokenString)

			json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
			return
		}
	}

	http.Error(w, "Wrong Username or Password.", http.StatusUnauthorized)
}

func AuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
			
		tokenString := strings.TrimPrefix(auth, "Bearer ")
		secret := os.Getenv("JWT_SECRET")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Forbiddent", http.StatusForbidden)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		user := models.User {
			ID:			int(claims["id"].(float64)),
			Username:	claims["username"].(string),
			Role:		claims["role"].(string),
		}

		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

