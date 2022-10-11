package gateway

import (
	"github.com/FANGDEI/simplebank/store/local"
	"github.com/jackc/pgconn"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type createAccountMessage struct {
	Owner    string `json:"owner" validate:"required"`
	Currency string `json:"currency" validate:"required,oneof=USD EUR CAD"`
}

type listAccountMessage struct {
	PageID   int `json:"page_id" validate:"required,min=1"`
	PageSize int `json:"page_size" validate:"required,min=5,max=10"`
}

func (m *Manager) RouteAccount() {
	m.handler.PartyFunc("/accounts", func(p iris.Party) {
		p.Post("/", m.createAccount)
		// iris在进行id范围校验时不成功则会返回404
		p.Get("/{id:int64 range(1, 9223372036854775807)}", m.getAccount)
		p.Post("/list", m.listAccount)
	})
}

func (m *Manager) createAccount(ctx iris.Context) {
	var msg createAccountMessage
	if err := ctx.ReadJSON(&msg); err != nil {
		m.sendSimpleMessage(ctx, iris.StatusBadRequest, err)
		return
	}

	errs := m.validator.Struct(msg)
	if errs != nil {
		m.sendValidateMessage(ctx, errs)
		return
	}

	account, err := m.localer.CreateAccount(local.Account{
		Owner:    msg.Owner,
		Balance:  0,
		Currency: msg.Currency,
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
		"data": account,
	})
}

func (m *Manager) getAccount(ctx iris.Context) {
	id, err := ctx.Params().GetInt64("id")
	if err != nil {
		m.sendSimpleMessage(ctx, iris.StatusBadRequest, err)
		return
	}

	account, err := m.localer.GetAccount(id)
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

	m.sendJson(ctx, iris.StatusOK, map[string]any{
		"data": account,
	})
}

func (m *Manager) listAccount(ctx iris.Context) {
	var msg listAccountMessage
	if err := ctx.ReadJSON(&msg); err != nil {
		m.sendSimpleMessage(ctx, iris.StatusBadRequest, err)
		return
	}

	errs := m.validator.Struct(msg)
	if errs != nil {
		m.sendValidateMessage(ctx, errs)
		return
	}

	accounts, err := m.localer.ListAccounts(msg.PageSize, (msg.PageID-1)*msg.PageSize)
	if err != nil {
		m.sendSimpleMessage(ctx, iris.StatusInternalServerError, err)
		return
	}

	m.sendJson(ctx, iris.StatusOK, map[string]any{
		"data": accounts,
	})
}
