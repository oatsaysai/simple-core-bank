package model

type TransferInParams struct {
	ToAccountNo string  `json:"to_account_no" validate:"required"`
	Amount      float64 `json:"amount" validate:"required,gt=0"`
}

type TransferInResponse struct {
	TransactionID int64   `json:"transaction_id"`
	ToAccountNo   string  `json:"to_account_no"`
	Amount        float64 `json:"amount"`
}

type TransferOutParams struct {
	FromAccountNo string  `json:"from_account_no" validate:"required"`
	Amount        float64 `json:"amount" validate:"required,gt=0"`
}

type TransferOutResponse struct {
	TransactionID int64   `json:"transaction_id"`
	FromAccountNo string  `json:"from_account_no"`
	Amount        float64 `json:"amount"`
}

type TransferParams struct {
	FromAccountNo string  `json:"from_account_no" validate:"required"`
	ToAccountNo   string  `json:"to_account_no" validate:"required"`
	Amount        float64 `json:"amount" validate:"required,gt=0"`
}

type TransferResponse struct {
	TransactionID int64   `json:"transaction_id"`
	FromAccountNo string  `json:"from_account_no"`
	ToAccountNo   string  `json:"to_account_no"`
	Amount        float64 `json:"amount"`
}

type TransferForLoadTestParams struct {
	MaxAccountNo int     `json:"max_account_no" validate:"required"`
	Amount       float64 `json:"amount" validate:"required,gt=0"`
}
