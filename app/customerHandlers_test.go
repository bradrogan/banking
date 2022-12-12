package app

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/bradrogan/banking/errs"
	"github.com/bradrogan/banking/mock"
	"github.com/golang/mock/gomock"
)

func TestCustomerHandlers_getAllCustomers(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name           string
		prepare        func(m *mock.MockCustomerServicer)
		args           args
		wantStatusCode int
	}{
		{
			name: "Test get all customers",
			prepare: func(m *mock.MockCustomerServicer) {
				m.EXPECT().GetAllCustomers().Return(nil, errs.NewNotFoundError("dummyErr"))
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/customers", nil),
			},
			wantStatusCode: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			service := mock.NewMockCustomerServicer(ctrl)
			if tt.prepare != nil {
				tt.prepare(service)
			}
			ch := &CustomerHandlers{
				service: service,
			}
			ch.getAllCustomers(tt.args.w, tt.args.r)

			got := tt.args.w.(*httptest.ResponseRecorder).Result().StatusCode

			if !reflect.DeepEqual(got, tt.wantStatusCode) {
				t.Errorf("CustomerHandlers.getallCustomers() = %v, want %v", got, tt.wantStatusCode)
			}

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
