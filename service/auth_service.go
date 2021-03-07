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

// Login  verifies a users credentails and then returns a jwt token on success
func (s DefaultAuthService) Login(req dto.LoginRequest) (*string, *errs.AppError) {
	login, err := s.repo.FindBy(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	token, err := login.GenerateToken()
	if err != nil {
		return nil, err
	}
	return token, nil
}
