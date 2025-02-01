package user

import (
	"context"
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetUserByEmail(ctx context.Context, email string) (User, error) {
	params := SelectOneUserInput{
		Email: &email,
	}
	return s.repo.SelectOneUser(ctx, params)
}

func (s *service) GetUserByID(ctx context.Context, id int64) (User, error) {
	params := SelectOneUserInput{
		UserID: &id,
	}
	return s.repo.SelectOneUser(ctx, params)
}

func (s *service) GetUsers(ctx context.Context) ([]User, error) {
	return s.repo.SelectUsers(ctx)
}

func (s *service) Register(ctx context.Context, input RegisterPayload) error {
	existingUser, err := s.GetUserByEmail(ctx, input.Email)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if existingUser.ID != 0 {
		return errors.New("user already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	input.Password = string(hashedPassword)
	return s.repo.InsertUser(ctx, input)
}
