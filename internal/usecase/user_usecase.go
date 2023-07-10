package usecase

import (
	"qvtec/go-app/internal/domain"
	"qvtec/go-app/internal/repository"
)

type UserUseCase interface {
	GetAll() ([]*domain.User, error)
	Create(user *domain.User) error
	GetByID(id int) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id int) error
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}

func (uc *userUseCase) GetAll() ([]*domain.User, error) {
	// @todo validation
	users, err := uc.userRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uc *userUseCase) Create(user *domain.User) error {
	// @todo validation
	err := uc.userRepository.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (uc *userUseCase) GetByID(id int) (*domain.User, error) {
	// @todo validation
	user, err := uc.userRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *userUseCase) Update(user *domain.User) error {
	// @todo validation
	err := uc.userRepository.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (uc *userUseCase) Delete(id int) error {
	// @todo validation
	err := uc.userRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
