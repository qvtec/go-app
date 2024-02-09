package usecase_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/qvtec/go-app/internal/domain"
	"github.com/qvtec/go-app/internal/usecase"
	"github.com/qvtec/go-app/mocks"
)

func setupMock() (*mocks.UserRepository, []*domain.User) {
	mockRepo := new(mocks.UserRepository)

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

	return mockRepo, mockUsers
}

func TestUserUseCase_GetAll(t *testing.T) {
	mockRepo, mockUsers := setupMock()

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("GetAll").Return(mockUsers, nil).Once()
		uc := usecase.NewUserUseCase(mockRepo)
		users, err := uc.GetAll()
		assert.NoError(t, err)
		assert.NotNil(t, users)
		assert.Len(t, users, len(mockUsers))
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		mockRepo.On("GetAll").Return(nil, errors.New("failed to fetch")).Once()
		uc := usecase.NewUserUseCase(mockRepo)
		users, err := uc.GetAll()
		assert.Error(t, err)
		assert.Nil(t, users)
		mockRepo.AssertExpectations(t)
	})
}

func TestUserUseCase_Create(t *testing.T) {
	mockRepo, mockUsers := setupMock()
	mockUser := mockUsers[0]

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("Create", mock.Anything).Return(nil).Once()
		uc := usecase.NewUserUseCase(mockRepo)
		err := uc.Create(mockUser)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		mockRepo.On("Create", mock.Anything).Return(errors.New("user not found")).Once()
		uc := usecase.NewUserUseCase(mockRepo)
		err := uc.Create(mockUser)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestUserUseCase_GetByID(t *testing.T) {
	mockRepo, mockUsers := setupMock()
	mockUser := mockUsers[0]

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("GetByID", mock.Anything).Return(mockUser, nil).Once()
		uc := usecase.NewUserUseCase(mockRepo)
		user, err := uc.GetByID(mockUser.ID)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, mockUser, user)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		mockRepo.On("GetByID", mock.Anything).Return(nil, errors.New("user not found")).Once()
		uc := usecase.NewUserUseCase(mockRepo)
		user, err := uc.GetByID(mockUser.ID)
		assert.Error(t, err)
		assert.Nil(t, user)
		mockRepo.AssertExpectations(t)
	})
}

func TestUserUseCase_Update(t *testing.T) {
	mockRepo, mockUsers := setupMock()
	mockUser := mockUsers[0]

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("Update", mock.Anything).Return(nil).Once()
		uc := usecase.NewUserUseCase(mockRepo)
		err := uc.Update(mockUser)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		mockRepo.On("Update", mock.Anything).Return(errors.New("user not found")).Once()
		uc := usecase.NewUserUseCase(mockRepo)
		err := uc.Update(mockUser)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestUserUseCase_Delete(t *testing.T) {
	mockRepo, mockUsers := setupMock()
	mockUser := mockUsers[0]

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("Delete", mock.Anything).Return(nil).Once()
		uc := usecase.NewUserUseCase(mockRepo)
		err := uc.Delete(mockUser.ID)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		mockRepo.On("Delete", mock.Anything).Return(errors.New("user not found")).Once()
		uc := usecase.NewUserUseCase(mockRepo)
		err := uc.Delete(mockUser.ID)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}