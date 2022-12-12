package app

import (
	"net/http"
	"reflect"
	"testing"
)

func TestAccountHandlers_NewAccount(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		ah   *AccountHandlers
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ah.NewAccount(tt.args.w, tt.args.r)
		})
	}
}

func TestAccountHandlers_NewTransaction(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		ah   *AccountHandlers
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ah.NewTransaction(tt.args.w, tt.args.r)
		})
	}
}

func TestNewAccountHandler(t *testing.T) {
	type args struct {
		a accountServicer
	}
	tests := []struct {
		name string
		args args
		want *AccountHandlers
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAccountHandler(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccountHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
