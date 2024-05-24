package controller

import (
	"encoding/json"
	"net/http"
	"oauth/model"
	"oauth/service"
)

// GetUserInfo handles GET requests to fetch user info.
func UserInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		userID := r.URL.Query().Get("id")
		user, err := service.GetUserInfo(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(user)
	}

	if r.Method == "POST" {
		var user model.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := service.AddUserInfo(user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}

}
