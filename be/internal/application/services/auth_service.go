package services

import (
	"fmt"

	"github.com/m4har/sams/internal/domain/dto"
	"github.com/m4har/sams/internal/domain/entity"
	"github.com/m4har/sams/internal/domain/repository"
	"github.com/m4har/sams/pkg/utils"
)

type AuthService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(email, password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &entity.User{
		Email:    email,
		Password: hashedPassword,
	}

	return s.userRepo.Create(user)
}

func (s *AuthService) Login(email, password string) (*dto.LoginResponseDTO, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, fmt.Errorf("invalid credentials")
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponseDTO{
		Token: token,
		User: dto.UserDTO{
			ID:        user.ID,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}, nil
}

func (s *AuthService) GetUserByID(id uint) (*entity.User, error) {
	return s.userRepo.FindByID(id)
}
