package app

import (
	"net/http"
	"testing"
)

func TestCustomerHandlers_getAllCustomers(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		ch   *CustomerHandlers
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ch.getAllCustomers(tt.args.w, tt.args.r)
		})
	}
}

func TestCustomerHandlers_getCustomersByStatus(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		ch   *CustomerHandlers
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ch.getCustomersByStatus(tt.args.w, tt.args.r)
		})
	}
}

func TestCustomerHandlers_getCustomer(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		ch   *CustomerHandlers
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ch.getCustomer(tt.args.w, tt.args.r)
		})
	}
}
