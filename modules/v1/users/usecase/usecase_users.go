package usecase

import (
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"master-user/helper"
	shareModel "master-user/model"
	"master-user/modules/v1/users/dto"
	"master-user/modules/v1/users/mapper"
	"master-user/modules/v1/users/model"
	"master-user/modules/v1/users/repository"

	"github.com/sirupsen/logrus"
)

type UserUsecaseImpl struct {
	userRepository repository.IUserRepository
	userMapper     mapper.IUserMapper
}

func NewUserUsecaseImpl(userRepository repository.IUserRepository, userMapper mapper.IUserMapper) IUserUsecase {
	return &UserUsecaseImpl{userRepository, userMapper}
}

func (u *UserUsecaseImpl) RegisterNewUser(request model.RegisterNewUserRequest) (user dto.UserDTO, err error) {
	ctx := "usecase_user.register_new_user"

	toModelUser := u.userMapper.ToUserModel(request)

	//Validasi Email
	isAvailable := u.userRepository.CheckEmail(toModelUser.Email)
	if !isAvailable {
		helper.Log(logrus.ErrorLevel, "email sudah terdaftar", ctx, "register new user")
		return user, errors.New("email sudah terdaftar")
	}

	registerUser := u.userRepository.Create(toModelUser)

	if registerUser.Error != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "register new user")
		return user, registerUser.Error
	}

	result := registerUser.Result.(shareModel.User)

	ud := u.userMapper.ToUserDTO(result)

	return ud, nil
}

func (u *UserUsecaseImpl) UpdateUser(userID uuid.UUID, request model.UpdateUserRequest) (user dto.UserDTO, err error) {
	ctx := "usecase_update_user.register_new_user"

	toModelUser := u.userMapper.ToUserModel(request)

	// Find User Id
	resultRepo := u.userRepository.FindById(userID)
	if resultRepo.Error != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "update user")
		return user, resultRepo.Error
	}

	result := resultRepo.Result.(shareModel.User)
	if result.Email != request.Email {
		//Validasi Email
		isAvailable := u.userRepository.CheckEmail(toModelUser.Email)
		if !isAvailable {
			helper.Log(logrus.ErrorLevel, "email sudah terdaftar", ctx, "update user")
			return user, errors.New("email sudah terdaftar")
		}
	}

	update := u.userRepository.Update(result.Id, toModelUser)
	if update.Error != nil {
		helper.Log(logrus.ErrorLevel, update.Error.Error(), ctx, "update user")
		return user, update.Error
	}
	finalResult := update.Result.(shareModel.User)
	ud := u.userMapper.ToUserDTO(finalResult)
	return ud, nil
}

func (u *UserUsecaseImpl) FindAllUsers() (users []dto.UserDTO, err error) {
	ctx := "usecase_update_user.register_new_user"

	resultRepository := u.userRepository.GetAll()
	if resultRepository.Error != nil {
		helper.Log(logrus.ErrorLevel, resultRepository.Error.Error(), ctx, "find all user")
		return users, resultRepository.Error
	}

	finalResult := resultRepository.Result.([]shareModel.User)
	userDTO := u.userMapper.ToListUserDTO(finalResult)

	return userDTO, nil
}

func (u *UserUsecaseImpl) DeleteUser(userID uuid.UUID) (message string, err error) {
	ctx := "usecase_update_user.delete_user"

	result := u.userRepository.Delete(userID)
	if result.Error != nil {
		helper.Log(logrus.ErrorLevel, result.Error.Error(), ctx, "delete user")
		return "", result.Error
	}

	finalResult := result.Result.(string)

	return finalResult, nil
}

func (u *UserUsecaseImpl) FindUserById(Id uuid.UUID) (user dto.UserDTO, err error) {
	ctx := "usecase_update_user.find_user_by_id"

	result := u.userRepository.FindById(Id)
	if result.Error != nil {
		helper.Log(logrus.ErrorLevel, result.Error.Error(), ctx, "find user by id")
		return user, result.Error
	}

	finalResult := result.Result.(shareModel.User)
	ud := u.userMapper.ToUserDTO(finalResult)
	return ud, nil
}
