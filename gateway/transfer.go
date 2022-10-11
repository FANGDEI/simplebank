package gateway

import (
	"errors"
	"fmt"

	"github.com/FANGDEI/simplebank/store/local"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type transferMessage struct {
	FromAccountID int64  `json:"from_account_id" validate:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" validate:"required,min=1"`
	Amount        int64  `json:"amount" validate:"required,gt=0"`
	Currency      string `json:"currency" validate:"required,oneof=USD EUR CAD"`
}

func (m *Manager) RouteTransfer() {
	m.handler.PartyFunc("/transfers", func(p iris.Party) {
		p.Use(m.tokener.Serve)
		p.Post("/", m.createTransfer)
	})
}

func (m *Manager) createTransfer(ctx iris.Context) {
	var msg transferMessage
	if err := ctx.ReadJSON(&msg); err != nil {
		m.sendSimpleMessage(ctx, iris.StatusBadRequest, err)
		return
	}

	fromAccount, valid := m.validAccount(ctx, msg.FromAccountID, msg.Currency)
	if !valid {
		return
	}

	if fromAccount.Owner != m.getUsername(ctx) {
		err := errors.New("from account doesn't belong to the authenticated user")
		m.sendSimpleMessage(ctx, iris.StatusUnauthorized, err)
		return
	}

	_, valid = m.validAccount(ctx, msg.FromAccountID, msg.Currency)
	if !valid {
		return
	}

	result, err := m.storer.TransferTx(local.TransferTxParams{
		FromAccountID: msg.FromAccountID,
		ToAccountID:   msg.ToAccountID,
		Amount:        msg.Amount,
	})
	if err != nil {
		m.sendSimpleMessage(ctx, iris.StatusInternalServerError, err)
		return
	}

	m.sendJson(ctx, iris.StatusOK, map[string]any{
		"data": result,
	})
}

func (m *Manager) validAccount(ctx iris.Context, accountID int64, currency string) (local.Account, bool) {
	account, err := m.localer.GetAccount(accountID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			m.sendJson(ctx, iris.StatusNotFound, map[string]any{
				"msg": err.Error(),
			})
			return account, false
		}
		m.sendSimpleMessage(ctx, iris.StatusInternalServerError, err)
		return account, false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", account.ID, account.Currency, currency)
		m.sendJson(ctx, iris.StatusBadRequest, map[string]any{
			"msg": err.Error(),
		})
		return account, false
	}
	return account, true
}
