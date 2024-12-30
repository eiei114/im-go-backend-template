package service

import (
	"context"
	"log"

	"github.com/eiei114/go-backend-template/domain"
	"github.com/eiei114/go-backend-template/domain/repository"
	"github.com/google/uuid"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (u *UserService) Add(ctx context.Context, name string) (string, error) {
	// UUIDでユーザIDと認証トークンを生成
	userID, err := uuid.NewRandom()
	if err != nil {
		log.Println("Failed to generate user ID", err)
		return "Failed to generate user ID", err
	}

	authToken, err := uuid.NewRandom()
	if err != nil {
		log.Println("Failed to generate auth token", err)
		return "Failed to generate auth token", err
	}

	err = u.UserRepository.AddUser(ctx, userID.String(), authToken.String(), name)
	if err != nil {
		log.Println("Failed to add user", err)
		return "Failed to add user", err
	}

	log.Println("User added successfully", userID.String(), authToken.String(), name)
	return authToken.String(), nil
}

func (u *UserService) UpdateUser(ctx context.Context, user domain.User) error {
	_ = u.UserRepository.UpdateUser(ctx, user)
	return nil
}

func (u *UserService) Delete(ctx context.Context, id string) (string, error) {
	_ = u.UserRepository.DeleteUser(ctx, id)
	return "", nil
}

func (u *UserService) GetUserByUserId(ctx context.Context, id string) (domain.User, error) {
	var user domain.User
	user, _ = u.UserRepository.GetUserByUserId(ctx, id)
	return user, nil
}

func (u *UserService) GetUserByAuthToken(ctx context.Context, authToken string) (*domain.User, error) {
	var user *domain.User
	user, _ = u.UserRepository.GetUserByAuthToken(ctx, authToken)
	return user, nil
}
