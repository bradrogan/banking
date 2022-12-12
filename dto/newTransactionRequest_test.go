package dto

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/bradrogan/banking/errs"
)

func TestNewTransactionRequst_Validate(t *testing.T) {
	tests := []struct {
		name string
		tr   NewTransactionRequst
		want *errs.AppError
	}{
		{
			name: "Test 'deposit' Transaction Type",
			tr: NewTransactionRequst{
				CustomerId: "1",
				AccountId:  "1",
				Type:       "deposit",
				Amount:     0.0,
			},
			want: nil,
		},
		{
			name: "Test 'withdrawal' Transaction Type",
			tr: NewTransactionRequst{
				CustomerId: "1",
				AccountId:  "1",
				Type:       "withdrawal",
				Amount:     0.0,
			},
			want: nil,
		},
		{
			name: "Test Invalid Transaction Type",
			tr: NewTransactionRequst{
				CustomerId: "1",
				AccountId:  "1",
				Type:       "",
				Amount:     0.0,
			},
			want: &errs.AppError{
				Code:    http.StatusUnprocessableEntity,
				Message: "invalid transaction type",
			},
		},
		{
			name: "Test Negative Amount",
			tr: NewTransactionRequst{
				CustomerId: "1",
				AccountId:  "1",
				Type:       "withdrawal",
				Amount:     -1,
			},
			want: &errs.AppError{
				Code:    http.StatusUnprocessableEntity,
				Message: "amount cannot be negative",
			},
		},
		{
			name: "Test Zero Amount",
			tr: NewTransactionRequst{
				CustomerId: "1",
				AccountId:  "1",
				Type:       "withdrawal",
				Amount:     0.0,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Validate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTransactionRequst.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTransactionRequst_IsWithdrawal(t *testing.T) {
	tests := []struct {
		name string
		tr   NewTransactionRequst
		want bool
	}{
		{
			name: "Test is withdrawal",
			tr: NewTransactionRequst{
				CustomerId: "",
				AccountId:  "",
				Type:       "withdrawal",
				Amount:     0.0,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.IsWithdrawal(); got != tt.want {
				t.Errorf("NewTransactionRequst.IsWithdrawal() = %v, want %v", got, tt.want)
			}
		})
	}
}
