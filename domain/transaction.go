package domain

type Transaction struct {
	Id        string
	AccountId string
	Type      TransactionType
	Amount    float64
	Time      string
}

type TransactionType string

const (
	Deposit    TransactionType = "deposit"
	Withdrawal TransactionType = "withdrawal"
)
