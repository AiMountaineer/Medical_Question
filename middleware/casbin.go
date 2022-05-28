package middleware

import (
	"MedicalQuestionnaireService/libs"
	"errors"
	"fmt"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/fatih/color"
	"github.com/kataras/iris/v12"
)

func New(e *casbin.Enforcer) *Casbin {
	return &Casbin{enforcer: e}
}

type Casbin struct {
	enforcer *casbin.Enforcer
}

func (c *Casbin) AuthToken(ctx iris.Context) {
	ctx.StatusCode(http.StatusOK)
	user := ctx.User()
	userid, _ := user.GetID()
	check, err := c.Check(ctx.Request(), userid)
	if !check {
		ctx.JSON(libs.ApiResource(403, nil, err.Error()))
		ctx.StopExecution()
		return
	} else {
		ctx.Values().Set("auth_user_id", userid)
	}
	ctx.Next()
}

// Check checks the username, request's method and path and
// returns true if permission grandted otherwise false.

func (c *Casbin) Check(r *http.Request, userId string) (bool, error) {
	method := r.Method
	path := r.URL.Path
	ok, err := c.enforcer.Enforce(userId, path, method)

	if err != nil {
		color.Red("验证权限报错：%v;%s-%s-%s", err.Error(), userId, path, method)
		return false, err
	}
	if !ok {
		return ok, errors.New(fmt.Sprintf("你未拥有 %s:%s 操作权限", method, path))
	}
	return ok, nil
}
