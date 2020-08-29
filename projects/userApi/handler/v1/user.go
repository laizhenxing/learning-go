package v1

import (
	"fmt"
	"userApi/pkg/auth"
	"userApi/pkg/token"

	//"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"github.com/unknwon/com"

	"userApi/handler"
	"userApi/model"
	"userApi/pkg/errno"
	"userApi/util"
)

type UserListResponse struct {
	Users []*model.User `json:"users"`
	Count uint64        `json:"count"`
}

func Create(c *gin.Context) {
	log.Info("User crate function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	var err error
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := model.User{
		Username: username,
		Password: password,
	}
	// 校验参数
	if err = user.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, handler.EmptyResponse)
		return
	}
	//valid := validation.Validation{}
	//valid.Required(username, "username").Message("名称不能为空")
	//valid.MaxSize(username, 20, "username").Message("名称长度不能超过20字符")
	//valid.Required(password, "password").Message("密码不能为空")
	//valid.MinSize(password, 4, "password").Message("密码长度不能小于4个字符")
	//valid.MaxSize(password, 16, "password").Message("密码长度不能超过16个字符")
	//if valid.HasErrors() {
	//	for _, err := range valid.Errors {
	//		log.Infof("Create user error: %v", err)
	//	}
	//	err = errno.New(*errno.ParamsError, err)
	//	handler.SendResponse(c, err, struct{}{})
	//	return
	//}

	// 密码加密
	if err = user.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, handler.EmptyResponse)
		return
	}

	// 插入数据库
	if err = user.Create(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, handler.EmptyResponse)
		return
	}

	handler.SendResponse(c, nil, map[string]string{"username": username})
}

// 更新用户
func Update(c *gin.Context) {
	userId := com.StrTo(c.Param("id")).MustInt()
	username := c.Query("username")
	password := c.Query("password")

	u := model.User{
		Username: username,
		Password: password,
	}
	u.ID = uint64(userId)

	// 校验参数
	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, handler.EmptyResponse)
		return
	}

	// 密码加密
	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, handler.EmptyResponse)
		return
	}

	if err := u.Update(); err != nil {
		fmt.Println(err)
		handler.SendResponse(c, errno.ErrDatabase, handler.EmptyResponse)
		return
	}

	handler.SendResponse(c, nil, handler.EmptyResponse)
}

// 删除用户
func Delete(c *gin.Context) {
	userId := com.StrTo(c.Param("id")).MustInt()

	err := model.DeleteUser(uint64(userId))
	if err != nil {
		handler.SendResponse(c, errno.ErrDatabase, handler.EmptyResponse)
		return
	}

	handler.SendResponse(c, nil, handler.EmptyResponse)
}

// 获取一个用户的信息
func Get(c *gin.Context) {
	username := c.Param("username")
	user, err := model.GetUser(username)
	if err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, handler.EmptyResponse)
		return
	}
	handler.SendResponse(c, nil, user)
}

// 获取多个用户的信息
func GetList(c *gin.Context) {
	limit := com.StrTo(c.Query("pageNum")).MustInt()
	offset := com.StrTo(c.Query("page")).MustInt()
	username := c.Query("username")

	users, count, err := model.GetUserList(username, offset, limit)
	if err != nil {
		handler.SendResponse(c, err, handler.EmptyResponse)
		return
	}

	handler.SendResponse(c, nil, UserListResponse{
		Users: users,
		Count: count,
	})

}

// 用户登录
func Login(c *gin.Context) {
	var err error
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := model.User{
		Username: username,
		Password: password,
	}
	// 校验参数
	if err = user.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, handler.EmptyResponse)
		return
	}

	// 校验用户是否存在
	u, err := model.GetUser(username)
	if err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, handler.EmptyResponse)
		return
	}

	// 检验密码
	if err = auth.Compare(u.Password, user.Password); err != nil {
		handler.SendResponse(c, errno.ErrPasswordIncorrect, handler.EmptyResponse)
		return
	}

	// 生成token
	tokenString, err := token.Sign(c, token.Claims{
		ID:       u.ID,
		Username: u.Username,
	}, "")
	if err != nil {
		handler.SendResponse(c, errno.ErrToken, handler.EmptyResponse)
		return
	}

	handler.SendResponse(c, nil, token.Token{Token: tokenString})
}
