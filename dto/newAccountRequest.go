package dto

type NewAccountRequest struct {
	Type       string  `json:"type,omitempty" xml:"type"`
	Amount     float64 `json:"amount,omitempty" xml:"amount"`
	CustomerId string  `json:"customer_id,omitempty" xml:"customer_id"`
}
