package postgres

import (
	"context"
	"testing"

	"github.com/raisa320/Labora-wallet/models"
)

func TestCreate(t *testing.T) {
	ctx := context.Background()
	var transaction models.Transaction = models.Transaction{
		Amount:    20,
		DestinyId: 9,
		SourceId:  10,
		Type:      2,
	}
	go func() {
		NewTransactionStorage().Create(ctx, transaction)
	}()

	go func() {
		NewTransactionStorage().Create(ctx, transaction)
	}()
}
