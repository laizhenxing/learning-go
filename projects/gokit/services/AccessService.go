package services

import (
	"errors"

	"gokit/util"
)

type IAccessService interface {
	GetToken(name, password string)	(string, error)
}

type AccessService struct {

}

func (a *AccessService) GetToken(name, password string) (string, error) {
	if name == "xingxiaoli" && password == "123" {
		return util.GenerateToken(name)
	}
	return "", errors.New("用户名或者密码不正确")
}




