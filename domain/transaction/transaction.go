package transaction

type Transaction struct {
	Id        string
	AccountId string
	Type      Type
	Amount    float64
	Time      string
}

type Type string

const (
	Deposit    Type = "deposit"
	Withdrawal Type = "withdrawal"
)
