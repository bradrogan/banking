package dto

import (
	"net/http"
	"reflect"
	"strconv"
	"testing"

	"github.com/bradrogan/banking/errs"
)

func TestNewAccountRequest_Validate(t *testing.T) {
	tests := []struct {
		name string
		r    NewAccountRequest
		want *errs.AppError
	}{
		{
			name: "Test invalid transaction type",
			r: NewAccountRequest{
				Type:       "",
				Amount:     5000,
				CustomerId: "",
			},
			want: &errs.AppError{
				Code:    http.StatusUnprocessableEntity,
				Message: "invalid account type",
			},
		},
		{
			name: "Test 'saving' transaction type",
			r: NewAccountRequest{
				Type:       "saving",
				Amount:     5000,
				CustomerId: "",
			},
			want: nil,
		},
		{
			name: "Test below minimum opening balance",
			r: NewAccountRequest{
				Type:       "saving",
				Amount:     MinimumOpeningBalance - 0.01,
				CustomerId: "",
			},
			want: &errs.AppError{
				Code:    http.StatusUnprocessableEntity,
				Message: "account opening requires a minimum $" + strconv.FormatFloat(MinimumOpeningBalance, 'f', 2, 64) + " deposit",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Validate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccountRequest.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
