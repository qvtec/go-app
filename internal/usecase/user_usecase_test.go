package usecase_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/qvtec/go-app/internal/domain"
	"github.com/qvtec/go-app/internal/usecase"
)

type mockUserRepository struct {
	mock.Mock
}

func (m *mockUserRepository) GetAll() ([]*domain.User, error) {
	args := m.Called()
	return args.Get(0).([]*domain.User), args.Error(1)
}

func (m *mockUserRepository) Create(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *mockUserRepository) GetByID(id int) (*domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *mockUserRepository) Update(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *mockUserRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestUserUseCase_GetAll(t *testing.T) {
	mockRepo := new(mockUserRepository)

	mockUsers := []*domain.User{
		{
			Name:  "John Doe",
			Email: "john@example.com",
		},
		{
			Name:  "Jane Smith",
			Email: "jane@example.com",
		},
	}
	mockRepo.On("GetAll").Return(mockUsers, nil)

	userUC := usecase.NewUserUseCase(mockRepo)

	// GetAll
	users, err := userUC.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, users)
	assert.Len(t, users, 2)
	assert.Equal(t, mockUsers, users)

	mockRepo.AssertExpectations(t)
}

func TestUserUseCase_Create(t *testing.T) {
	mockRepo := new(mockUserRepository)

	user := &domain.User{
		Name:  "John Doe",
		Email: "john@example.com",
	}
	mockRepo.On("Create", user).Return(nil)

	userUC := usecase.NewUserUseCase(mockRepo)

	// Create
	err := userUC.Create(user)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestUserUseCase_GetByID(t *testing.T) {
	mockRepo := new(mockUserRepository)

	mockUser := &domain.User{
		Name:  "John Doe",
		Email: "john@example.com",
	}
	mockRepo.On("GetByID", mock.AnythingOfType("int")).Return(mockUser, nil)
	mockRepo.On("GetByID", mock.AnythingOfType("int")).Return(nil, errors.New("user not found"))

	userUC := usecase.NewUserUseCase(mockRepo)

	// GetByID（存在するユーザの場合）
	fetchedUser, err := userUC.GetByID(mockUser.ID)
	assert.NoError(t, err)
	assert.NotNil(t, fetchedUser)
	assert.Equal(t, mockUser, fetchedUser)

	// GetByID（存在しないユーザの場合）
	// fetchedUser, err = userUC.GetByID(100)
	// assert.Error(t, err)
	// assert.Nil(t, fetchedUser)

	mockRepo.AssertExpectations(t)
}

func TestUserUseCase_Update(t *testing.T) {
	mockRepo := new(mockUserRepository)

	user := &domain.User{
		Name:  "John Doe",
		Email: "john@example.com",
	}
	mockRepo.On("Update", user).Return(nil)

	userUC := usecase.NewUserUseCase(mockRepo)

	// Update
	err := userUC.Update(user)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestUserUseCase_Delete(t *testing.T) {
	mockRepo := new(mockUserRepository)

	mockRepo.On("Delete", mock.AnythingOfType("int")).Return(nil)

	userUC := usecase.NewUserUseCase(mockRepo)

	// Delete
	err := userUC.Delete(1)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
