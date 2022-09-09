package presenter

import (
	uuid "github.com/satori/go.uuid"
	"master-user/helper"
	"master-user/modules/v1/users/model"
	"master-user/modules/v1/users/usecase"
	"net/http"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type HttpHandlerUser struct {
	userUsecase usecase.IUserUsecase
}

func NewHttpHandlerUser(userUsecase usecase.IUserUsecase) *HttpHandlerUser {
	return &HttpHandlerUser{userUsecase}
}

func (h *HttpHandlerUser) Mount(group *echo.Group) {
	group.POST("/user", h.RegisterNewUser)
	group.PUT("/user/:id", h.UpdateUser)
	group.GET("/user", h.FindAllUsers)
	group.DELETE("/user/:id", h.DeleteUser)
	group.GET("/user/:id", h.FindUserById)
}

func (h *HttpHandlerUser) RegisterNewUser(e echo.Context) (err error) {
	ctx := "user_handler.register_new_user"

	userRegisterRequest := new(model.RegisterNewUserRequest)
	err = e.Bind(&userRegisterRequest)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "failed binding for Insert data user")
		return e.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err.Error(), nil))
	}

	result, err := h.userUsecase.RegisterNewUser(*userRegisterRequest)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "failed Insert data user")
		return e.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err.Error(), nil))
	}

	return e.JSON(http.StatusCreated, helper.ResponseDetailOutput("success register new user", result))
}

func (h *HttpHandlerUser) UpdateUser(e echo.Context) (err error) {
	ctx := "user_handler.update_user"

	Id, _ := uuid.FromString(e.Param("id"))

	updateUser := new(model.UpdateUserRequest)
	err = e.Bind(&updateUser)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "failed binding for update data user")
		return e.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err.Error(), nil))
	}

	result, err := h.userUsecase.UpdateUser(Id, *updateUser)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "failed update data user")
		return e.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err.Error(), nil))
	}

	return e.JSON(http.StatusCreated, helper.ResponseDetailOutput("success update user", result))
}

func (h *HttpHandlerUser) FindAllUsers(e echo.Context) (err error) {
	ctx := "user_handler.update_user"

	users, err := h.userUsecase.FindAllUsers()
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "find all users")
		return e.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err.Error(), nil))
	}

	return e.JSON(http.StatusCreated, helper.ResponseDetailOutput("success load users", users))
}

func (h *HttpHandlerUser) DeleteUser(e echo.Context) (err error) {
	ctx := "user_handler.delete_user"

	Id, _ := uuid.FromString(e.Param("id"))
	user, err := h.userUsecase.DeleteUser(Id)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "delete user")
		return e.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err.Error(), nil))
	}

	return e.JSON(http.StatusCreated, helper.ResponseDetailOutput(user, nil))
}

func (h *HttpHandlerUser) FindUserById(e echo.Context) (err error) {
	ctx := "user_handler.find_user_by_id"

	Id, _ := uuid.FromString(e.Param("id"))
	user, err := h.userUsecase.FindUserById(Id)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "find user by id")
		return e.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err.Error(), nil))
	}

	return e.JSON(http.StatusCreated, helper.ResponseDetailOutput("success load user by id", user))
}
