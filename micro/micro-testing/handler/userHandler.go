package handler

import (
	"context"

	"micro-testing/model"
	"micro-testing/proto/user"
)

type UserHandler struct {

}

func (u UserHandler) UserReg(ctx context.Context, user *user.UserModel, rsp *user.UserResponse) error {
	usrData := model.User{
		UserId:   int(user.UserId),
		UserName: user.UserName,
		Password: user.UserPassword,
	}
	if err := model.GetDB().Create(&usrData).Error; err != nil {
		rsp.Message = err.Error()
		rsp.Status = "Error"
	} else {
		rsp.Message = "create user success~"
		rsp.Status = "Success"
	}
	return nil
}
