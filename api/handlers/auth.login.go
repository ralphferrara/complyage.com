package handlers

import (
	"api/db"
	"api/errors"
	"api/models"
	"api/responses"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var input LoginInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		responses.Error(w, http.StatusBadRequest, errors.ErrInvalidRequest)
		return
	}

	pepper := os.Getenv("PASSWORD_PEPPER")
	fmt.Printf("PEPPER: [%s]\n", pepper)
	fmt.Printf("Login attempt: email=%s password=%s\n", input.Email, input.Password)

	if pepper == "" {
		responses.Error(w, http.StatusInternalServerError, errors.ErrMissingPepper)
		return
	}

	user := getUserByEmail(input.Email)
	if user == nil {
		responses.Error(w, http.StatusUnauthorized, errors.ErrInvalidCredentials)
		return
	}

	passwordWithPepper := input.Password + pepper

	fmt.Println([]byte(user.UserPassword))
	fmt.Println([]byte(passwordWithPepper))

	err := bcrypt.CompareHashAndPassword([]byte(user.UserPassword), []byte(passwordWithPepper))
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, errors.ErrInvalidCredentials)
		return
	}

	responses.Success(w, http.StatusOK, map[string]any{
		"user_id": user.IDUser,
		"email":   user.UserEmail,
	})
}

func getUserByEmail(email string) *models.User {
	var user models.User

	result := db.DB.Where("user_email = ?", email).First(&user)
	if result.Error != nil {
		return nil
	}

	return &user
}
