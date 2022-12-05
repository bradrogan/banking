package dto

type NewTransactionResponse struct {
	TransactionId   string  `json:"transaction_id,omitempty" xml:"transaction_id"`
	TransactionType string  `json:"transaction_type,omitempty" xml:"transaction_type"`
	AccountId       string  `json:"account_id,omitempty" xml:"account_id"`
	Balance         float64 `json:"balance,omitempty" xml:"balance"`
	TransactionDate string  `json:"transaction_date,omitempty" xml:"transaction_date"`
}
