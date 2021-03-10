package service

import (
	"github.com/jonathanwamsley/banking_auth/domain"
	"github.com/jonathanwamsley/banking_auth/dto"
	"github.com/jonathanwamsley/banking_auth/errs"
)

type UsersService interface {
	CreateUser(user dto.UserRequest) (*dto.UserResponse, *errs.AppError)
	CreateAdmin(admin dto.AdminRequest) (*dto.AdminResponse, *errs.AppError)
	GetUsers() ([]dto.UserResponse, *errs.AppError)
}

type DefaultUsersService struct {
	repo domain.UsersRepository
}

func NewUsersService(repo domain.UsersRepository) DefaultUsersService {
	return DefaultUsersService{repo}
}

func (s DefaultUsersService) CreateUser(user dto.UserRequest) (*dto.UserResponse, *errs.AppError) {

	err := user.Validate()
	if err != nil {
		return nil, err
	}
	newUser := domain.NewUser(user)
	result, err := s.repo.CreateUser(newUser)
	if err != nil {
		return nil, err
	}
	resp := domain.UserResponse(*result)
	return &resp, nil
}

func (s DefaultUsersService) CreateAdmin(admin dto.AdminRequest) (*dto.AdminResponse, *errs.AppError) {
	err := admin.Validate()
	if err != nil {
		return nil, err
	}
	newAdmin := domain.NewAdmin(admin)
	result, err := s.repo.CreateAdmin(newAdmin)
	if err != nil {
		return nil, err
	}
	resp := domain.AdminResponse(*result)
	return &resp, nil
}

func (s DefaultUsersService) GetUsers() ([]dto.UserResponse, *errs.AppError) {
	users, err := s.repo.GetUsers()
	if err != nil {
		return nil, err
	}
	usersResponse := make([]dto.UserResponse, 0)
	for _, user := range users {
		usersResponse = append(usersResponse, domain.UserResponse(user))
	}
	return usersResponse, nil
}
