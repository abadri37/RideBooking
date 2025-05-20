package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"ridebooking/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// --- Mock UserService ---

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) RegisterUser(ctx context.Context, req model.UserRequest) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *MockUserService) LoginUser(ctx context.Context, email, password string) (string, error) {
	args := m.Called(ctx, email, password)
	return args.String(0), args.Error(1)
}

func (m *MockUserService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(*model.User), args.Error(1)
}
func (m *MockUserService) GetUserByUserId(ctx context.Context, id string) (*model.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserService) UpdateUser(ctx context.Context, req model.UserRequest) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *MockUserService) RemoveUser(ctx context.Context, email string) error {
	args := m.Called(ctx, email)
	return args.Error(0)
}

// --- Tests ---

func TestRegisterUserHandler(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewUserHandler(mockService)

	userReq := model.UserRequest{
		Email:     "test@example.com",
		Password:  "secret",
		FirstName: "Test",
		LastName:  "User",
		Type:      "Rider",
		UserId:    "U123",
	}
	body, _ := json.Marshal(userReq)

	mockService.On("RegisterUser", mock.Anything, userReq).Return(nil)

	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	handler.RegisterUserHandler(rr, req)

	require.Equal(t, http.StatusCreated, rr.Code)
	assert.JSONEq(t, `{"message":"User registered successfully"}`, rr.Body.String())
	mockService.AssertExpectations(t)
}

func TestUserLoginHandler_Success(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewUserHandler(mockService)

	loginReq := model.LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}
	body, _ := json.Marshal(loginReq)

	mockService.On("LoginUser", mock.Anything, loginReq.Email, loginReq.Password).Return("mock-token", nil)

	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	handler.UserLoginHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `{"Token":"mock-token"}`, rr.Body.String())
	mockService.AssertExpectations(t)
}

func TestGetUserByEmailHandler(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewUserHandler(mockService)

	mockUser := &model.User{Email: "test@example.com", FirstName: "Test"}
	mockService.On("GetUserByEmail", mock.Anything, "test@example.com").Return(mockUser, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/ridebooking/user/emailId?emailId=test@example.com", nil)
	rr := httptest.NewRecorder()

	handler.GetUserByEmailHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), `"Email":"test@example.com"`)
}

func TestGetUserByIdHandler(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewUserHandler(mockService)

	mockUser := &model.User{UserId: "U123", FirstName: "John"}
	mockService.On("GetUserByUserId", mock.Anything, "U123").Return(mockUser, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/ridebooking/user/id?userId=U123", nil)
	rr := httptest.NewRecorder()

	handler.GetUserByIdHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), `"UserId":"U123"`)
}

func TestUpdateUserHandler(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewUserHandler(mockService)

	userReq := model.UserRequest{
		Email:     "updated@example.com",
		FirstName: "Updated",
	}
	body, _ := json.Marshal(userReq)

	mockService.On("UpdateUser", mock.Anything, userReq).Return(nil)

	req := httptest.NewRequest(http.MethodPut, "/api/ridebooking/user", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	handler.UpdateUserHandler(rr, req)

	assert.Equal(t, http.StatusAccepted, rr.Code)
	assert.Contains(t, rr.Body.String(), "User Updated successfully")
}

func TestRemoveUserByEmailHandler(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewUserHandler(mockService)

	mockService.On("RemoveUser", mock.Anything, "test@example.com").Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/api/ridebooking/user/emailId?emailId=test@example.com", nil)
	rr := httptest.NewRecorder()

	handler.RemoveUserByEmailHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "User Removed successfully")
}
