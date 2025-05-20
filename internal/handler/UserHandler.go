package handler

import (
	"encoding/json"
	"net/http"
	_ "ridebooking/docs"
	"ridebooking/internal/model"
	"ridebooking/internal/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// @Summary Register a new user
// @Description Registers a new user in the system
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.UserRequest true "User registration details"
// @Success 201 {object} map[string]string
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Internal server error"
// @Router /register [post]
func (h *UserHandler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var userRequest model.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := h.userService.RegisterUser(r.Context(), userRequest); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

// @Summary User login
// @Description Authenticates user and returns JWT token
// @Tags users
// @Accept json
// @Produce json
// @Param credentials body model.LoginRequest true "User credentials"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Invalid JSON payload"
// @Failure 500 {string} string "Authentication failed"
// @Router /login [post]
func (h *UserHandler) UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	var req model.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Authenticate and get token
	token, err := h.userService.LoginUser(r.Context(), req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.LoginResponse{Token: token})
}

// @Summary Get user by email
// @Description Retrieves user details by email ID
// @Tags users
// @Accept json
// @Produce json
// @Param emailId query string true "Email ID of the user"
// @Success 200 {object} model.User
// @Failure 500 {string} string "Internal server error"
// @Router /api/ridebooking/user/emailId [get]
func (h *UserHandler) GetUserByEmailHandler(w http.ResponseWriter, r *http.Request) {

	emailId := r.URL.Query().Get("emailId")
	user, err := h.userService.GetUserByEmail(r.Context(), emailId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// @Summary Get user by ID
// @Description Retrieves user details by user ID
// @Tags users
// @Accept json
// @Produce json
// @Param userId query string true "User ID"
// @Success 200 {object} model.User
// @Failure 500 {string} string "Internal server error"
// @Router /api/ridebooking/user/id [get]
func (h *UserHandler) GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {

	emailId := r.URL.Query().Get("userId")
	user, err := h.userService.GetUserByUserId(r.Context(), emailId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// @Summary Update user details
// @Description Updates an existing user's details
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.UserRequest true "Updated user info"
// @Success 202 {object} map[string]string
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Internal server error"
// @Router /api/ridebooking/user [put]
func (h *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var userRequest model.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := h.userService.UpdateUser(r.Context(), userRequest); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"message": "User Updated successfully"})
}

// @Summary Delete user by email
// @Description Deletes a user based on email ID
// @Tags users
// @Accept json
// @Produce json
// @Param emailId query string true "Email ID"
// @Success 200 {object} map[string]string
// @Failure 500 {string} string "Internal server error"
// @Router /api/ridebooking/user/emailId [delete]
func (h *UserHandler) RemoveUserByEmailHandler(w http.ResponseWriter, r *http.Request) {

	emailId := r.URL.Query().Get("emailId")
	err := h.userService.RemoveUser(r.Context(), emailId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User Removed successfully"})
}
