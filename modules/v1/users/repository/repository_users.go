package repository

import (
	"errors"
	"master-user/helper"
	"master-user/model"
	"master-user/modules/share"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type UserRepositoryImpl struct {
	*share.Repository
}

func NewUserRepositoryImpl(repo *share.Repository) IUserRepository {
	return &UserRepositoryImpl{repo}
}

func (r *UserRepositoryImpl) Create(user model.User) share.ResultRepository {
	ctx := "repository_users.create"

	_, err := r.DBWrite.Model(&user).Insert()
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, " create user")
		return share.ResultRepository{Error: err}
	}

	return share.ResultRepository{Result: user, Error: nil}
}

func (r *UserRepositoryImpl) Update(userID uuid.UUID, user model.User) share.ResultRepository {
	ctx := "repository_users.update"

	tx, err := r.DBWrite.Begin()
	defer tx.Close()

	if err != nil {
		tx.Rollback()
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "Cannot begin transaction in update user")
		return share.ResultRepository{Error: err}
	}

	res, err := tx.Model(&user).
		Where("id = ? ", userID).
		UpdateNotZero()
	if err != nil {
		tx.Rollback()
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "Cannot update user in transaction")
		return share.ResultRepository{Error: err}
	}
	if res.RowsAffected() <= 0 {
		tx.Rollback()
		err = errors.New("User with specified ID does not exist or already delete")
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "Cannot update data user")
		return share.ResultRepository{Error: err}
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "Cannot commit transaction update data user")
		return share.ResultRepository{Error: err}
	}

	resultUser := r.FindById(userID).Result
	return share.ResultRepository{Error: nil, Result: resultUser}

}

func (r *UserRepositoryImpl) Delete(userId uuid.UUID) share.ResultRepository {
	ctx := "repository_users.delete"

	user := model.User{IsActive: false, UpdatedAt: time.Now()}
	tx, err := r.DBWrite.Begin()
	defer tx.Close()
	if err != nil {
		tx.Rollback()
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "Cannot begin transaction in delete user")
		return share.ResultRepository{Error: err}
	}
	// Delete User Group
	res, err := tx.Query(&user, `
		UPDATE public.users SET is_active = false, updated_at = ?
		WHERE id = ? AND is_active = true
	`, time.Now(), userId)

	if err != nil {
		tx.Rollback()
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "Cannot delete data user")
		return share.ResultRepository{Error: err}
	}

	if res.RowsAffected() <= 0 {
		tx.Rollback()
		err = errors.New("User with specified ID does not exist or already delete")
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "Cannot delete data user")
		return share.ResultRepository{Error: err}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "Cannot commit transaction data user")
		return share.ResultRepository{Error: err}
	}

	successMessage := "User deleted successfully"
	return share.ResultRepository{Result: successMessage}
}

func (r *UserRepositoryImpl) GetAll() share.ResultRepository {
	ctx := "repository_user.get_all"

	users := []model.User{}

	err := r.DBRead.Model(&users).Where("is_active = ? ", true).Select()
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "get all")
		return share.ResultRepository{Error: err}
	}

	return share.ResultRepository{Result: users}
}

func (r *UserRepositoryImpl) FindById(userId uuid.UUID) share.ResultRepository {
	ctx := "repository_user.find_by_id"

	user := model.User{}

	err := r.DBRead.Model(&user).Where("id = ? AND is_active = true ", userId).Select()
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "find by id")
		return share.ResultRepository{Error: err}
	}

	return share.ResultRepository{Result: user}
}

func (r *UserRepositoryImpl) CheckEmail(email string) (isAvailable bool) {
	ctx := "repository_user.check_email"

	user := model.User{}

	err := r.DBRead.Model(&user).Where("email = ? ", email).Select()
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "check email")
		return true
	}

	if user.Email == "" {
		return true
	}

	return false
}
