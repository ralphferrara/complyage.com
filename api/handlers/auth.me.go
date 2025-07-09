package handlers

import (
	"api/responses"
	"net/http"
)

func AuthMeHandler(w http.ResponseWriter, r *http.Request) {
	user, err := getUserFromSession(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	responses.Success(w, http.StatusOK, map[string]any{
		"user_id": user.IDUser,
		"email":   user.UserEmail,
	})
}
