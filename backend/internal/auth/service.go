package auth

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Register(
	ctx context.Context,
	request RegisterRequest,
) (*RegisterResponse, error) {

	existingUser, _ := s.repository.FindUserByEmail(
		ctx,
		request.Email,
	)

	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, err
	}

	user := User{
		ID:           uuid.New().String(),
		Name:         request.Name,
		Email:        request.Email,
		PasswordHash: string(passwordHash),
		Role:         "student",
		TenantID:     request.TenantID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err = s.repository.CreateUser(ctx, user)

	if err != nil {
		return nil, err
	}

	return &RegisterResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Role:     user.Role,
		TenantID: user.TenantID,
	}, nil
}

func (s *Service) Login(
	ctx context.Context,
	request LoginRequest,
	secret string,
) (*LoginResponse, error) {
	user, err := s.repository.FindUserByEmail(
		ctx,
		request.Email,
	)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(request.Password),
	)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}
	token, err := GenerateToken(
		*user,
		secret,
	)
	if err != nil {
		return nil, err
	}
	return &LoginResponse{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   24 * 3600,
	}, nil
}
