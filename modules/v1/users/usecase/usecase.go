package usecase

import (
	uuid "github.com/satori/go.uuid"
	"master-user/modules/v1/users/dto"
	"master-user/modules/v1/users/model"
)

type IUserUsecase interface {
	RegisterNewUser(request model.RegisterNewUserRequest) (user dto.UserDTO, err error)
	UpdateUser(userID uuid.UUID, request model.UpdateUserRequest) (user dto.UserDTO, err error)
	FindAllUsers() (users []dto.UserDTO, err error)
	DeleteUser(userID uuid.UUID) (message string, err error)
	FindUserById(Id uuid.UUID) (user dto.UserDTO, err error)
}
