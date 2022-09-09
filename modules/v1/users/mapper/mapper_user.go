package mapper

import (
	"master-user/helper"
	"master-user/model"
	"master-user/modules/v1/users/dto"
	shareModel "master-user/modules/v1/users/model"
	"time"

	uuid "github.com/satori/go.uuid"
)

type UserMapperImpl struct {
}

func NewUserMapperImpl() IUserMapper {
	return &UserMapperImpl{}
}

func (m *UserMapperImpl) ToUserModel(request any) (user model.User) {

	switch request.(type) {
	case shareModel.RegisterNewUserRequest:
		data := request.(shareModel.RegisterNewUserRequest)
		passwordHash, _ := helper.HashPassword(data.Password)
		uuidV1 := uuid.NewV1()

		user.Id = uuidV1
		user.Nama = data.Nama
		user.Email = data.Email
		user.Username = data.Username
		user.Password = passwordHash
		user.IsActive = true
		user.CreatedAt = time.Now()

	case shareModel.UpdateUserRequest:
		data := request.(shareModel.UpdateUserRequest)

		user.Nama = data.Nama
		user.Email = data.Email
		user.Username = data.Username
		user.UpdatedAt = time.Now()

	}
	return user
}

func (m *UserMapperImpl) ToUserDTO(data model.User) (user dto.UserDTO) {
	return dto.UserDTO{
		Nama:      data.Nama,
		Email:     data.Email,
		Username:  data.Username,
		IsActive:  data.IsActive,
		CreatedAt: data.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func (m *UserMapperImpl) ToListUserDTO(list []model.User) (users []dto.UserDTO) {
	for _, V := range list {
		users = append(users, dto.UserDTO{
			Nama:      V.Nama,
			Email:     V.Email,
			Username:  V.Username,
			IsActive:  V.IsActive,
			CreatedAt: V.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: V.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return users
}
