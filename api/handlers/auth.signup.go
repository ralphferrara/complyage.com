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

type SignupInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var input SignupInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		responses.Error(w, http.StatusBadRequest, errors.ErrInvalidRequest)
		return
	}

	pepper := os.Getenv("PASSWORD_PEPPER")
	if pepper == "" {
		responses.Error(w, http.StatusInternalServerError, errors.ErrMissingPepper)
		return
	}

	// Check if user already exists
	var existing models.User
	if err := db.DB.Where("user_email = ?", input.Email).First(&existing).Error; err == nil {
		responses.Error(w, http.StatusBadRequest, "user_already_exists")
		return
	}

	// Hash password + pepper
	combined := input.Password + pepper
	hash, err := bcrypt.GenerateFromPassword([]byte(combined), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Hashing error:", err)
		responses.Error(w, http.StatusInternalServerError, "hashing_failed")
		return
	}

	newUser := models.User{
		UserEmail:    input.Email,
		UserPassword: string(hash),
		UserStatus:   "ACTV",
		UserMerchant: false,
	}

	if err := db.DB.Create(&newUser).Error; err != nil {
		fmt.Println("DB create error:", err)
		responses.Error(w, http.StatusInternalServerError, "db_insert_failed")
		return
	}

	responses.Success(w, http.StatusOK, map[string]any{
		"user_id": newUser.IDUser,
		"email":   newUser.UserEmail,
	})
}
