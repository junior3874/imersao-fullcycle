package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID           string
	SellingOrder *Order
	BuyingOrder  *Order
	Shares       int
	Price        float64
	Total        float64
	DateTime     time.Time
}

func NewTransaction(sellingOrder *Order, buyingOrder *Order, shares int, price float64) *Transaction {
	total := float64(shares) * price

	return &Transaction{
		ID:           uuid.New().String(),
		SellingOrder: sellingOrder,
		BuyingOrder:  buyingOrder,
		Shares:       shares,
		Price:        price,
		Total:        total,
		DateTime:     time.Now(),
	}
}

func (t *Transaction) GetMinShares() int {
	sellingShares := t.SellingOrder.PendingShares
	buyingShare := t.BuyingOrder.PendingShares

	if sellingShares > buyingShare {
		return sellingShares
	} else {
		return buyingShare
	}
}

func (t *Transaction) HasExistPendingSharesInSellingOrder() bool {
	return t.SellingOrder.PendingShares > 0
}

func (t *Transaction) HasExistPendingSharesInBuyingOrder() bool {
	return t.BuyingOrder.PendingShares > 0
}

func (t *Transaction) CloseSellingOrder() {
	t.SellingOrder.Status = "CLOSED"
}

func (t *Transaction) CloseBuyinhOrder() {
	t.BuyingOrder.Status = "CLOSED"
}

func (t *Transaction) SetTotalTransactionPrice() {
	t.Total = float64(t.Shares) * t.BuyingOrder.Price
}

func (t *Transaction) MakeTransaction() {
	minsShares := t.GetMinShares()
	t.SellingOrder.Investor.UpdateAssetPosition(t.SellingOrder.Asset.ID, -minsShares)
	t.SellingOrder.ChangeShareValue(minsShares)
	t.BuyingOrder.Investor.UpdateAssetPosition(t.SellingOrder.Asset.ID, minsShares)
	t.BuyingOrder.ChangeShareValue(minsShares)
	t.SetTotalTransactionPrice()

	if t.HasExistPendingSharesInBuyingOrder() {
		t.CloseBuyinhOrder()
	}

	if t.HasExistPendingSharesInSellingOrder() {
		t.CloseSellingOrder()
	}

}
