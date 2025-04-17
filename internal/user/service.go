package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Signup(input SignupInput) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

type SignupInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (s *service) Signup(input SignupInput) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	if err := s.repo.Create(user); err != nil {
		return errors.New("could not create user: " + err.Error())
	}

	return nil
}
