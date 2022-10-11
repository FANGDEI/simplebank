package gateway

import (
	"github.com/FANGDEI/simplebank/store/local"
	"github.com/jackc/pgconn"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type createUserMessage struct {
	Username string `json:"username" validate:"required,alphanum"`
	Password string `json:"password" validate:"required,min=6"`
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

type loginMessage struct {
	Username string `json:"username" validate:"required,alphanum"`
	Password string `json:"password" validate:"required,min=6"`
}

func (m *Manager) RouteUser() {
	m.handler.PartyFunc("/users", func(p iris.Party) {
		p.Post("/", m.createUser)
		p.Post("/login", m.loginUser)
	})
}

func (m *Manager) createUser(ctx iris.Context) {
	var msg createUserMessage
	if err := ctx.ReadJSON(&msg); err != nil {
		m.sendSimpleMessage(ctx, iris.StatusBadRequest, err)
		return
	}

	errs := m.validator.Struct(msg)
	if errs != nil {
		m.sendValidateMessage(ctx, errs)
		return
	}

	user, err := m.localer.CreateUser(local.SimpleUser{
		Username: msg.Username,
		Password: m.cryptoer.ToMd5(msg.Password),
		FullName: msg.FullName,
		Email:    msg.Email,
	})
	if err != nil {
		if _, ok := err.(*pgconn.PgError); ok {
			m.sendSimpleMessage(ctx, iris.StatusBadRequest, err)
			return
		}
		m.sendSimpleMessage(ctx, iris.StatusInternalServerError, err)
		return
	}

	m.sendJson(ctx, iris.StatusOK, map[string]any{
		"data": user,
	})
}

func (m *Manager) loginUser(ctx iris.Context) {
	var msg loginMessage
	if err := ctx.ReadJSON(&msg); err != nil {
		m.sendSimpleMessage(ctx, iris.StatusBadRequest, err)
		return
	}

	errs := m.validator.Struct(msg)
	if errs != nil {
		m.sendValidateMessage(ctx, errs)
		return
	}

	user, err := m.localer.GetUser(msg.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			m.sendJson(ctx, iris.StatusNotFound, map[string]any{
				"msg": err.Error(),
			})
			return
		}
		m.sendSimpleMessage(ctx, iris.StatusInternalServerError, err)
		return
	}

	if m.cryptoer.ToMd5(msg.Password) != user.Password {
		m.sendSimpleMessage(ctx, iris.StatusUnauthorized, err)
		return
	}

	token, err := m.newToken(user.Username)
	if err != nil {
		m.sendSimpleMessage(ctx, iris.StatusInternalServerError, err)
		return
	}

	m.sendJson(ctx, iris.StatusOK, map[string]any{
		"token": token,
		"data":  user,
	})
}
