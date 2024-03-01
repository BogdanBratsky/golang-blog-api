package handlers

import (
	"encoding/json"
	"golang-blog-api/internal/services"
	"golang-blog-api/models"
	"log"
	"net/http"
)

type UserHandler struct {
	AuthService services.AuthService
}

func NewUserHandler(authService services.AuthService) *UserHandler {
	return &UserHandler{
		AuthService: authService,
	}
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Request recieved: ", r.Method, r.URL)

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("Failed to decode JSON request: ", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	passwordHashed, err := h.AuthService.PasswordHash(user.UserPassword)
	if err != nil {
		log.Println("Failed to hash password: ", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	user.UserPassword = passwordHashed

	response := models.Response{
		Success: true,
		Message: "sdfcsd",
		Data:    user,
		// Data: models.UserResponse{
		// 	UserId:    user.UserId,
		// 	UserName:  user.UserName,
		// 	UserEmail: user.UserEmail,
		// 	CreatedAt: user.CreatedAt,
		// },
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		response := models.BadResponse{
			Success: false,
			Message: "Внутренняя ошибка сервера",
		}
		json.NewEncoder(w).Encode(response)
		log.Println("Failed to encode JSON: ", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
}

// func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Request recieved: ", r.Method, r.URL)

// 	var user models.User
// 	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
// 		log.Println("Failed to decode JSON: ", err)
// 		http.Error(w, "Bad request", http.StatusBadRequest)
// 		return
// 	}

// 	hashedPassword, err := services.PasswordHash(user.UserPassword)
// 	if err != nil {
// 		log.Println("Failed to hash password: ", err)
// 		http.Error(w, "Server error", http.StatusInternalServerError)
// 		return
// 	}
// 	user.UserPassword = hashedPassword

// 	response := struct {
// 		Success bool        `json:"success"`
// 		Message string      `json:"message"`
// 		Token   string      `json:"token"`
// 		User    models.User `json:"user"`
// 	}{
// 		Success: true,
// 		Message: "Пользователь успешно авторизован",
// 		Token:   "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
// 		User:    user,
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	if err := json.NewEncoder(w).Encode(response); err != nil {
// 		log.Println("Failed to encode JSON: ", err)
// 		http.Error(w, "Server error", http.StatusInternalServerError)
// 		return
// 	}
// }
