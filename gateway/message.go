package gateway

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

// var (
// 	unknownError = "未知错误, 请联系管理员或稍后重试"
// )

func (m *Manager) sendJson(ctx iris.Context, code int, v any) {
	ctx.StatusCode(code)
	ctx.JSON(v)
}

func (m *Manager) sendSimpleMessage(ctx iris.Context, code int, errs ...error) {
	if errs == nil {
		m.sendJson(ctx, code, map[string]string{
			"msg": "成功",
		})
		return
	}
	for _, err := range errs {
		log.Printf("[%s] CODE(%d) ERROR : %+v\n", ctx.Path(), code, err)
	}
	var msg string
	switch code {
	case iris.StatusBadRequest:
		msg = "请求解析失败"
	case iris.StatusInternalServerError:
		msg = "服务器内部错误，请联系管理员"
	case iris.StatusOK:
		msg = "请求成功"
	case iris.StatusPreconditionFailed:
		msg = "预处理失败"
	case iris.StatusForbidden:
		msg = "权限认证失败"
	default:
		msg = "未知错误"
	}
	m.sendJson(ctx, code, map[string]string{
		"msg": msg,
	})
}

func (m *Manager) sendValidateMessage(ctx iris.Context, err error) {
	for _, err := range err.(validator.ValidationErrors) {
		log.Printf("[%s] CODE(%d) ERROR : %+v\n", ctx.Path(), iris.StatusOK, err)
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(map[string]any{
		"msg": "数据校验失败",
	})
}
