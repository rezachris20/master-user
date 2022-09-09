package mapper

import (
	"master-user/model"
	"master-user/modules/v1/users/dto"
)

type IUserMapper interface {
	ToUserModel(request any) (user model.User)
	ToUserDTO(data model.User) (user dto.UserDTO)
	ToListUserDTO(list []model.User) (users []dto.UserDTO)
}
