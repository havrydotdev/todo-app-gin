package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/gavrylenkoIvan/todo-app-gin"
	"github.com/gavrylenkoIvan/todo-app-gin/pkg/repository"
)

const salt = "UUj9HKhoD90hLdh0DFHklgvasa"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
