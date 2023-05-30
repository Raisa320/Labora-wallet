package models

type Transaction struct {
	Id            int     `json:"id"`
	Amount        float64 `json:"amount" validate:"required"`
	WalletReceive int     `json:"walletReceive_id" validate:"required"`
	WalletOrigin  int     `json:"walletOrigin_id" validate:"required"`
	Type          string  `json:"type" validate:"required"`
}
