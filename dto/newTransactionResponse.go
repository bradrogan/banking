package dto

type NewTransactionResponse struct {
	Id      string  `json:"id,omitempty" xml:"id"`
	Balance float64 `json:"balance,omitempty" xml:"balance"`
}
