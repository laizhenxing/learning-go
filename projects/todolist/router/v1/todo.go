package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"todolist/model"
	"todolist/pkg/errno"
	"todolist/router"
)

func AddTodo(c *gin.Context)  {
	title := c.PostForm("title")

	valid := validation.Validation{}
	valid.Required(title, "title").Message("标题不能为空")

	if valid.HasErrors() {
		router.SendResponse(c, errno.ErrParams, router.EmptyResult())
		return
	}

	tl := model.TodoList{
		Title:  title,
		Status: false,
	}

	err := tl.CreateTodo()
	if err != nil {
		router.SendResponse(c, errno.NewErr(*errno.ErrDatabase, err), router.EmptyResult())
		return
	}

	router.SendResponse(c, nil, tl)
}

func Update(c *gin.Context)  {
	id := com.StrTo(c.Param("id")).MustInt()
	title := c.Query("title")
	status := com.StrTo(c.Query("status"))

	tl := model.TodoList{
		Model:  model.Model{
			ID: uint(id),
		},
		Title:  title,
		Status: status.Exist(),
	}
	if err := tl.UpdateTodo(); err != nil {
		router.SendResponse(c, errno.NewErr(*errno.ErrDatabase, err), router.EmptyResult())
		return
	}

	router.SendResponse(c, errno.OK, tl)
}

func GetTodo(c *gin.Context)  {
	id := com.StrTo(c.Param("id")).MustInt()

	if exist := model.ExistTodoByID(uint(id)); !exist {
		router.SendResponse(c, errno.ErrNotFound, router.EmptyResult())
		return
	}

	todo, err := model.GetTodo(uint(id))
	if err != nil {
		router.SendResponse(c, errno.NewErr(*errno.ErrDatabase, err), router.EmptyResult())
		return
	}

	router.SendResponse(c, errno.OK, todo)
}

func GetTodoList(c *gin.Context)  {
	title := c.Query("title")
	pageNum := com.StrTo(c.Query("pageNum")).MustInt()
	pageSize := com.StrTo(c.Query("pageSize")).MustInt()

	tls, err := model.GetTodoList(title, pageNum, pageSize)
	if err != nil {
		router.SendResponse(c, errno.NewErr(*errno.ErrDatabase, err), router.EmptyResult())
		return
	}

	router.SendResponse(c, errno.OK, tls)
}

func DeleteTodo(c *gin.Context)  {
	id := com.StrTo(c.Param("id")).MustInt()

	if exist := model.ExistTodoByID(uint(id)); !exist {
		router.SendResponse(c, errno.ErrNotFound, router.EmptyResult())
		return
	}

	if err := model.DeleteTodo(uint(id)); err != nil {
		router.SendResponse(c, errno.NewErr(*errno.ErrDatabase, err), router.EmptyResult())
		return
	}

	router.SendResponse(c, errno.OK, router.EmptyResult())
}
