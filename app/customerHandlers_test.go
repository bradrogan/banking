package app

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/bradrogan/banking/domain/customer"
	"github.com/bradrogan/banking/dto"
	"github.com/bradrogan/banking/errs"
	"github.com/bradrogan/banking/mock"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

func TestCustomerHandlers_getAllCustomers(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name             string
		setUp            func(m *mock.MockCustomerServicer)
		args             args
		wantStatusCode   int
		wantResponseBody string
	}{
		{
			name: "Test get all customers customers found",
			setUp: func(m *mock.MockCustomerServicer) {
				m.EXPECT().GetAllCustomers().Return(
					[]dto.CustomerResponse{{
						Id:          "1",
						Name:        "Test",
						City:        "City",
						Zipcode:     "90210",
						DateOfBirth: "01/01/2001",
						Status:      "active",
					}}, nil)
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/customers", nil),
			},
			wantStatusCode:   http.StatusOK,
			wantResponseBody: `[{"id":"1","name":"Test","city":"City","zipcode":"90210","date_of_birth":"01/01/2001","status":"active"}]`,
		},
		{
			name: "Test get all customers not found",
			setUp: func(m *mock.MockCustomerServicer) {
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
			if tt.setUp != nil {
				tt.setUp(service)
			}
			ch := &CustomerHandlers{
				service: service,
			}
			ch.getAllCustomers(tt.args.w, tt.args.r)

			resp := tt.args.w.(*httptest.ResponseRecorder).Result()
			body := parseResponse(resp)
			status := resp.StatusCode

			if !reflect.DeepEqual(status, tt.wantStatusCode) {
				t.Errorf("CustomerHandlers.getAllCustomers() = %v, want %v", status, tt.wantStatusCode)
			}

			if tt.wantResponseBody != "" && !reflect.DeepEqual(body, tt.wantResponseBody) {
				t.Errorf("CustomerHandlers.getAllCustomers() = %v, want %v", body, tt.wantResponseBody)
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
		name             string
		setUp            func(m *mock.MockCustomerServicer)
		args             args
		params           map[string]string
		wantStatusCode   int
		wantResponseBody string
	}{
		{
			name: "Test get active customers",
			setUp: func(m *mock.MockCustomerServicer) {
				m.EXPECT().GetCustomersByStatus(customer.Active).Return(
					[]dto.CustomerResponse{{
						Id:          "1",
						Name:        "Test",
						City:        "City",
						Zipcode:     "90210",
						DateOfBirth: "01/01/2001",
						Status:      "active",
					}}, nil)
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/customers", nil),
			},
			params:           map[string]string{"status": "active"},
			wantStatusCode:   http.StatusOK,
			wantResponseBody: `[{"id":"1","name":"Test","city":"City","zipcode":"90210","date_of_birth":"01/01/2001","status":"active"}]`,
		},

		{
			name: "Test get inactive customers",
			setUp: func(m *mock.MockCustomerServicer) {
				m.EXPECT().GetCustomersByStatus(customer.Inactive).Return(
					[]dto.CustomerResponse{}, nil)
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/customers", nil),
			},
			params:           map[string]string{"status": "inactive"},
			wantStatusCode:   http.StatusOK,
			wantResponseBody: `[]`,
		},

		{
			name: "Test get invalid status parameter",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/customers", nil),
			},
			params:           map[string]string{"status": ""},
			wantStatusCode:   http.StatusBadRequest,
			wantResponseBody: `{"message":"invalid value for 'status' query parameter"}`,
		},
		{
			name: "Test Customer Service error returned",
			setUp: func(m *mock.MockCustomerServicer) {
				m.EXPECT().GetCustomersByStatus(customer.Active).Return(
					nil,
					errs.NewUnexpectedError("dummy err"),
				)
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/customers", nil),
			},
			params:         map[string]string{"status": "active"},
			wantStatusCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		ctrl := gomock.NewController(t)
		service := mock.NewMockCustomerServicer(ctrl)
		if tt.setUp != nil {
			tt.setUp(service)
		}
		if tt.params != nil {
			tt.args.r = mux.SetURLVars(tt.args.r, tt.params)
		}
		ch := &CustomerHandlers{
			service: service,
		}
		ch.getCustomersByStatus(tt.args.w, tt.args.r)

		resp := tt.args.w.(*httptest.ResponseRecorder).Result()
		body := parseResponse(resp)
		status := resp.StatusCode

		if !reflect.DeepEqual(status, tt.wantStatusCode) {
			t.Errorf("CustomerHandlers.getCustomersByStatus() = %v, want %v", status, tt.wantStatusCode)
		}

		if tt.wantResponseBody != "" && !reflect.DeepEqual(body, tt.wantResponseBody) {
			t.Errorf("CustomerHandlers.getCustomersByStatus() = %v, want %v", body, tt.wantResponseBody)
		}
	}
}

func TestCustomerHandlers_getCustomer(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name             string
		setUp            func(m *mock.MockCustomerServicer)
		args             args
		params           map[string]string
		wantStatusCode   int
		wantResponseBody string
	}{
		{
			name: "Customer by Id is found",
			setUp: func(m *mock.MockCustomerServicer) {
				m.EXPECT().GetCustomer("1").Return(
					&dto.CustomerResponse{
						Id:          "1",
						Name:        "Test",
						City:        "City",
						Zipcode:     "90210",
						DateOfBirth: "01/01/2001",
						Status:      "active",
					}, nil)
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/customer", nil),
			},
			params:           map[string]string{"customer_id": "1"},
			wantStatusCode:   http.StatusOK,
			wantResponseBody: `{"id":"1","name":"Test","city":"City","zipcode":"90210","date_of_birth":"01/01/2001","status":"active"}`,
		},
		{
			name: "Customer by Id is not found",
			setUp: func(m *mock.MockCustomerServicer) {
				m.EXPECT().GetCustomer("1").Return(
					nil, errs.NewNotFoundError("dummy err"))
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/customer", nil),
			},
			params:         map[string]string{"customer_id": "1"},
			wantStatusCode: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		ctrl := gomock.NewController(t)
		service := mock.NewMockCustomerServicer(ctrl)
		if tt.setUp != nil {
			tt.setUp(service)
		}
		if tt.params != nil {
			tt.args.r = mux.SetURLVars(tt.args.r, tt.params)
		}
		ch := &CustomerHandlers{
			service: service,
		}
		ch.getCustomer(tt.args.w, tt.args.r)

		resp := tt.args.w.(*httptest.ResponseRecorder).Result()
		body := parseResponse(resp)
		status := resp.StatusCode

		if !reflect.DeepEqual(status, tt.wantStatusCode) {
			t.Errorf("CustomerHandlers.getCustomer() = %v, want %v", status, tt.wantStatusCode)
		}

		if tt.wantResponseBody != "" && !reflect.DeepEqual(body, tt.wantResponseBody) {
			t.Errorf("CustomerHandlers.getCustomer() = %v, want %v", body, tt.wantResponseBody)
		}
	}
}

func parseResponse(resp *http.Response) string {
	bodyBuffer := new(bytes.Buffer)
	bodyBuffer.ReadFrom(resp.Body)
	body := strings.TrimSpace(bodyBuffer.String())

	return body
}
