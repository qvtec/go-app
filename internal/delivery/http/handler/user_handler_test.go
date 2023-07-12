package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/qvtec/go-app/internal/delivery/http/handler"
	"github.com/qvtec/go-app/internal/domain"
)

type mockUserUseCase struct {
	mock.Mock
}

func (m *mockUserUseCase) GetAll() ([]*domain.User, error) {
	args := m.Called()
	return args.Get(0).([]*domain.User), args.Error(1)
}

func (m *mockUserUseCase) Create(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *mockUserUseCase) GetByID(id int) (*domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *mockUserUseCase) Update(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *mockUserUseCase) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestUserHandler_GetAll(t *testing.T) {
	mockUseCase := new(mockUserUseCase)

	mockUsers := []*domain.User{
		{
			ID:    1,
			Name:  "John Doe",
			Email: "john@example.com",
		},
		{
			ID:    2,
			Name:  "Jane Smith",
			Email: "jane@example.com",
		},
	}
	mockUseCase.On("GetAll").Return(mockUsers, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	userHandler := handler.NewUserHandler(mockUseCase)

	// GetAll
	userHandler.GetAll(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "john@example.com")
	assert.Contains(t, w.Body.String(), "jane@example.com")
	// assert.JSONEq(t, `[{"ID":1,"Name":"John Doe","Email":"john@example.com"},{"ID":2,"Name":"Jane Smith","Email":"jane@example.com"}]`, w.Body.String())

	mockUseCase.AssertExpectations(t)
}

func TestUserHandler_Create(t *testing.T) {
	mockUseCase := new(mockUserUseCase)

	mockUser := &domain.User{
		Name:  "John Doe",
		Email: "john@example.com",
	}
	mockUseCase.On("Create", mockUser).Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"Name":"John Doe","Email":"john@example.com"}`))
	c.Request.Header.Set("Content-Type", "application/json")

	userHandler := handler.NewUserHandler(mockUseCase)

	// Create
	userHandler.Create(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.JSONEq(t, `{"message":"User created successfully"}`, w.Body.String())

	mockUseCase.AssertExpectations(t)
}

func TestUserHandler_Get(t *testing.T) {
	mockUseCase := new(mockUserUseCase)

	mockUser := &domain.User{
		ID:    1,
		Name:  "John Doe",
		Email: "john@example.com",
	}
	mockUseCase.On("GetByID", mockUser.ID).Return(mockUser, nil)
	mockUseCase.On("GetByID", 100).Return(nil, errors.New("user not found"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	userHandler := handler.NewUserHandler(mockUseCase)

	// Get
	userHandler.Get(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "john@example.com")
	// assert.JSONEq(t, `{"user":{"ID":1,"Name":"John Doe","Email":"john@example.com"}}`, w.Body.String())

	// w = httptest.NewRecorder()
	// c, _ = gin.CreateTestContext(w)
	// c.Params = []gin.Param{{Key: "id", Value: "100"}}

	// userHandler.Get(c)

	// assert.Equal(t, http.StatusNotFound, w.Code)
	// assert.JSONEq(t, `{"error":"User not found"}`, w.Body.String())

	// mockUseCase.AssertExpectations(t)
}

func TestUserHandler_Update(t *testing.T) {
	mockUseCase := new(mockUserUseCase)

	mockUser := &domain.User{
		ID:    1,
		Name:  "John Doe",
		Email: "john@example.com",
	}
	mockUseCase.On("Update", mockUser).Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(`{"Name":"John Doe","Email":"john@example.com"}`))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	userHandler := handler.NewUserHandler(mockUseCase)

	// Update
	userHandler.Update(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message":"User updated successfully"}`, w.Body.String())
}

func TestUserHandler_Delete(t *testing.T) {
	mockUseCase := new(mockUserUseCase)

	mockUseCase.On("Delete", 1).Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	userHandler := handler.NewUserHandler(mockUseCase)

	// Delete
	userHandler.Delete(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message":"User deleted successfully"}`, w.Body.String())

	mockUseCase.AssertExpectations(t)
}
