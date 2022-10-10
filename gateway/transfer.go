package gateway

import (
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
	m.handler.Post("/transfers", m.createTransfer)
}

func (m *Manager) createTransfer(ctx iris.Context) {
	var msg transferMessage
	if err := ctx.ReadJSON(&msg); err != nil {
		m.sendSimpleMessage(ctx, iris.StatusBadRequest, err)
		return
	}

	if !m.validAccount(ctx, msg.FromAccountID, msg.Currency) {
		return
	}

	if !m.validAccount(ctx, msg.ToAccountID, msg.Currency) {
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

func (m *Manager) validAccount(ctx iris.Context, accountID int64, currency string) bool {
	account, err := m.localer.GetAccount(accountID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			m.sendJson(ctx, iris.StatusNotFound, map[string]any{
				"msg": err.Error(),
			})
			return false
		}
		m.sendSimpleMessage(ctx, iris.StatusInternalServerError, err)
		return false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", account.ID, account.Currency, currency)
		m.sendJson(ctx, iris.StatusBadRequest, map[string]any{
			"msg": err.Error(),
		})
		return false
	}
	return true
}
