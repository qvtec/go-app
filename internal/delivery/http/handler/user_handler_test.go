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
	"github.com/qvtec/go-app/mocks"
)

func setupMock() (*mocks.UserUseCase, []*domain.User) {
	mockUseCase := new(mocks.UserUseCase)

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

	return mockUseCase, mockUsers
}

func TestUserHandler_GetAll(t *testing.T) {
	mockUseCase, mockUsers := setupMock()

	t.Run("Success", func(t *testing.T) {
		mockUseCase.On("GetAll").Return(mockUsers, nil).Once()
		userHandler := handler.NewUserHandler(mockUseCase)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		userHandler.GetAll(c)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), `"id":1,"name":"John Doe","email":"john@example.com"`)
		assert.Contains(t, w.Body.String(), `"id":2,"name":"Jane Smith","email":"jane@example.com"`)
	})

	t.Run("Error", func(t *testing.T) {
		mockUseCase.On("GetAll").Return(nil, errors.New("error")).Once()
		userHandler := handler.NewUserHandler(mockUseCase)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		userHandler.GetAll(c)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		expectedResponse := `{"error": "Failed to get users"}`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})
}

func TestUserHandler_Create(t *testing.T) {
	mockUseCase, _ := setupMock()

	requestBody := `{"id": 1, "name": "Test User", "email": "test@test.com", "password": "password"}`
	reader := strings.NewReader(requestBody)

	t.Run("Success", func(t *testing.T) {
		mockUseCase.On("Create", mock.Anything).Return(nil).Once()
		userHandler := handler.NewUserHandler(mockUseCase)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/users", reader)
		c.Request.Header.Set("Content-Type", "application/json")

		userHandler.Create(c)
		assert.Equal(t, http.StatusCreated, w.Code)
		expectedResponse := `{"message": "User created successfully"}`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})

	t.Run("BadRequest", func(t *testing.T) {
		userHandler := handler.NewUserHandler(mockUseCase)

		requestBodyInvalid := `Invalid Request Body`
		reader := strings.NewReader(requestBodyInvalid)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/users", reader)
		c.Request.Header.Set("Content-Type", "application/json")

		userHandler.Create(c)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Error", func(t *testing.T) {
		mockUseCase.On("Create", mock.Anything).Return(errors.New("error")).Once()
		userHandler := handler.NewUserHandler(mockUseCase)

		reader := strings.NewReader(requestBody)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/users", reader)
		c.Request.Header.Set("Content-Type", "application/json")

		userHandler.Create(c)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		expectedResponse := `{"error": "Failed to create user"}`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})
}

func TestUserHandler_Get(t *testing.T) {
	mockUseCase, mockUsers := setupMock()
	mockUser := mockUsers[0]

	t.Run("Success", func(t *testing.T) {
		mockUseCase.On("GetByID", mock.Anything).Return(mockUser, nil).Once()
		userHandler := handler.NewUserHandler(mockUseCase)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = []gin.Param{{Key: "id", Value: "1"}}

		userHandler.GetByID(c)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("BadRequest", func(t *testing.T) {
		userHandler := handler.NewUserHandler(mockUseCase)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = []gin.Param{{Key: "id", Value: "invalid"}}

		userHandler.GetByID(c)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		// {"error":"Invalid ID"}
	})

	t.Run("UserNotFound", func(t *testing.T) {
		mockUseCase.On("GetByID", mock.Anything).Return(nil, nil).Once()
		userHandler := handler.NewUserHandler(mockUseCase)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = []gin.Param{{Key: "id", Value: "100"}}

		userHandler.GetByID(c)
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.JSONEq(t, `{"error": "User not found"}`, w.Body.String())
	})

	t.Run("Error", func(t *testing.T) {
		mockUseCase.On("GetByID", mock.Anything).Return(nil, errors.New("error")).Once()
		userHandler := handler.NewUserHandler(mockUseCase)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = []gin.Param{{Key: "id", Value: "100"}}

		userHandler.GetByID(c)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.JSONEq(t, `{"error": "Failed to get user"}`, w.Body.String())
	})
}

func TestUserHandler_Update(t *testing.T) {
	mockUseCase, _ := setupMock()

	requestBody := `{"name": "Test User", "email": "test@test.com"}`
	reader := strings.NewReader(requestBody)

	t.Run("Success", func(t *testing.T) {
		mockUseCase.On("Update", mock.Anything).Return(nil).Once()
		userHandler := handler.NewUserHandler(mockUseCase)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPut, "/users/1", reader)
		c.Request.Header.Set("Content-Type", "application/json")
		c.Params = []gin.Param{{Key: "id", Value: "1"}}

		userHandler.Update(c)
		assert.Equal(t, http.StatusOK, w.Code)
		expectedResponse := `{"message":"User updated successfully"}`
		assert.Equal(t, expectedResponse, w.Body.String())
	})

	t.Run("BadRequestId", func(t *testing.T) {
		mockUseCase.On("Update", mock.Anything).Return(nil).Once()
		userHandler := handler.NewUserHandler(mockUseCase)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPut, "/users/1", reader)
		c.Request.Header.Set("Content-Type", "application/json")
		c.Params = []gin.Param{{Key: "id", Value: "Invalid"}}

		userHandler.Update(c)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		expectedResponse := `{"error": "Invalid user ID"}`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})

	t.Run("BadRequest", func(t *testing.T) {
		mockUseCase.On("Update", mock.Anything).Return(errors.New("error")).Once()
		userHandler := handler.NewUserHandler(mockUseCase)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = []gin.Param{{Key: "id", Value: "1"}}

		userHandler.Update(c)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		expectedResponse := `{"error": "Invalid request body"}`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})
}

func TestUserHandler_Delete(t *testing.T) {
	mockUseCase, _ := setupMock()

	t.Run("Success", func(t *testing.T) {
		mockUseCase.On("Delete", 1).Return(nil).Once()
		userHandler := handler.NewUserHandler(mockUseCase)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = []gin.Param{{Key: "id", Value: "1"}}

		userHandler.Delete(c)
		assert.Equal(t, http.StatusOK, w.Code)
		expectedResponse := `{"message": "User deleted successfully"}`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})

	t.Run("BadRequest", func(t *testing.T) {
		mockUseCase.On("Delete", mock.Anything).Return(errors.New("error")).Once()
		userHandler := handler.NewUserHandler(mockUseCase)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = []gin.Param{{Key: "Invalid", Value: "1"}}

		userHandler.Delete(c)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"error":"Invalid user ID"}`, w.Body.String())
	})

	t.Run("Error", func(t *testing.T) {
		mockUseCase.On("Delete", mock.Anything).Return(errors.New("error")).Once()
		userHandler := handler.NewUserHandler(mockUseCase)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = []gin.Param{{Key: "id", Value: "1"}}

		userHandler.Delete(c)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.JSONEq(t, `{"error": "Failed to delete user"}`, w.Body.String())
	})
}
