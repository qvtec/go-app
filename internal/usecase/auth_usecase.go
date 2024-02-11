package usecase

import (
	"os"
	"time"

	"github.com/qvtec/go-app/internal/repository"
	"github.com/qvtec/go-app/pkg/crypto"
	"github.com/qvtec/go-app/pkg/jwt"
)

type AuthUseCase interface {
	Login(email string, password string) (string, error)
}

type authUseCase struct {
	authRepository repository.AuthRepository
}

func NewAuthUseCase(authRepository repository.AuthRepository) AuthUseCase {
	return &authUseCase{
		authRepository: authRepository,
	}
}

func (uc *authUseCase) Login(email, password string) (string, error) {
	user, err := uc.authRepository.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if err := crypto.CheckPasswordHash(password, user.Password); err != nil {
		return "", err
	}

	jwtManager := jwt.NewJWTManager(os.Getenv("JWT_KEY"))

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	tokenString, err := jwtManager.GenerateToken(claims)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
