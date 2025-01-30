package service

import (
	dto "github.com/richieieie/event-booking/internal/DTO"
	"github.com/richieieie/event-booking/internal/model"
	"github.com/richieieie/event-booking/internal/repository"
	"github.com/richieieie/event-booking/internal/utils"
)

type IUserService interface {
	SignUp(userDto dto.SignUpUserDTO) (int64, error)
	Login(userDto dto.LoginUserDTO) (string, error)
}

type userService struct {
	iUserRepository repository.IUserRepository
}

func NewUserService(iUserRepository repository.IUserRepository) *userService {
	// Add business logic here
	// For example: maybe we want to map to DTO here, or filter some events by a condition by default
	// I do not have any ideas right now, so I just put the response from repository call here
	return &userService{iUserRepository: iUserRepository}
}

func (s userService) Login(userDto dto.LoginUserDTO) (string, error) {
	user, err := s.iUserRepository.GetByEmailAndPassword(userDto.Email)
	if err != nil {
		return "", err
	}

	err = utils.CheckHashedPassword(user.Password, userDto.Password)
	if err != nil {
		return "", err
	}

	accessToken, err := utils.GenerateJwtToken(user)
	return accessToken, err
}

func (s userService) SignUp(userDto dto.SignUpUserDTO) (int64, error) {
	hashedPassword, err := utils.HashPassword(userDto.Password)
	if err != nil {
		return 0, err
	}

	return s.iUserRepository.CreateOne(model.User{
		Email:    userDto.Email,
		Password: hashedPassword,
	})
}
