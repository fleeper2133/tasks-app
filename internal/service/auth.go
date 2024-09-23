package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/fleeper2133/tasks-app/internal/domain"
	"github.com/fleeper2133/tasks-app/internal/pkg"
	"github.com/fleeper2133/tasks-app/internal/repository"
	"github.com/go-playground/validator/v10"
)

const (
	salt = "dsadsadsad321#dsa"
)

type AuthorizationService struct {
	repo       repository.Authorization
	jwtManager pkg.TokenJWTManager
}

func NewAuthorizationService(repo repository.Authorization, jwtManager pkg.TokenJWTManager) *AuthorizationService {
	return &AuthorizationService{repo: repo, jwtManager: jwtManager}
}

func (s *AuthorizationService) CreateUser(user domain.SignUp) (int, error) {
	v := validator.New()
	if err := v.Struct(user); err != nil {
		return 0, err
	}
	if user.Password != user.RetryPassword {
		return 0, fmt.Errorf("passwords must match")
	}
	user.Password = s.hashPassword(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthorizationService) hashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthorizationService) GenerateTokens(input domain.SignIn) (pkg.TokenJWT, error) {
	input.Password = s.hashPassword(input.Password)
	user, err := s.repo.GetUser(input)
	if err != nil {
		return pkg.TokenJWT{}, err
	}
	userId := fmt.Sprint(user.Id)
	return s.jwtManager.NewJWTtoken(userId)
}

func (s *AuthorizationService) ParseToken(token string) (string, error) {
	return s.jwtManager.ParseAccessToken(token)
}

func (s *AuthorizationService) RefreshToken(refreshToken string) (pkg.TokenJWT, error) {
	return s.jwtManager.RefreshToken(refreshToken)
}
