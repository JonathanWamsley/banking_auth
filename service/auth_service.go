package service

import (
	"github.com/jonathanwamsley/banking_auth/domain"
	"github.com/jonathanwamsley/banking_auth/dto"
	"github.com/jonathanwamsley/banking_auth/errs"
)

type AuthService interface {
	Login(dto.LoginRequest) (*string, *errs.AppError)
}

type DefaultAuthService struct {
	repo domain.AuthRepository
}

func NewLoginService(repo domain.AuthRepository) DefaultAuthService {
	return DefaultAuthService{repo}
}

// Login  verifies a users credentails and then returns dummy msg on sucess
func (s DefaultAuthService) Login(req dto.LoginRequest) (*string, *errs.AppError) {
	_, err := s.repo.FindBy(req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	msg := "mock successfully logged in"
	return &msg, nil
}
