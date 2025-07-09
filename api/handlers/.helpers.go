package handlers

import (
	"api/db"
	"api/models"
	"errors"
	"net/http"
	"strconv"
)

func getUserFromSession(r *http.Request) (*models.User, error) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return nil, errors.New("no session")
	}

	sessionID := cookie.Value

	// Example: Parse your session_id to get the user ID.
	// ✅ In production, you'd store session_id → user_id in Redis or DB.
	userID, err := strconv.ParseUint(sessionID, 10, 64)
	if err != nil {
		return nil, errors.New("invalid session")
	}

	var user models.User
	result := db.DB.First(&user, userID)
	if result.Error != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}
