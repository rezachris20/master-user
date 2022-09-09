package repository

import (
	"master-user/model"
	"master-user/modules/share"

	uuid "github.com/satori/go.uuid"
)

type IUserRepository interface {
	Create(user model.User) share.ResultRepository
	Update(userID uuid.UUID, user model.User) share.ResultRepository
	Delete(userId uuid.UUID) share.ResultRepository
	GetAll() share.ResultRepository
	FindById(userId uuid.UUID) share.ResultRepository
	CheckEmail(email string) (isAvailable bool)
}
