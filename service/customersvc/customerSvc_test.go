package customersvc

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/bradrogan/banking/domain/customer"
	"github.com/bradrogan/banking/dto"
	"github.com/bradrogan/banking/errs"
	"github.com/bradrogan/banking/mock"
	"github.com/golang/mock/gomock"
)

func TestCustomerService_GetAllCustomers(t *testing.T) {
	tests := []struct {
		name  string
		setUp func(*mock.MockcustomerData)
		want  []dto.CustomerResponse
		want1 *errs.AppError
	}{
		{
			name: "Test got customers",
			setUp: func(m *mock.MockcustomerData) {
				m.EXPECT().FindAll().Return(
					[]customer.Customer{{
						Id:          "1",
						Name:        "Bob",
						City:        "Boston",
						Zipcode:     "90210",
						DateOfBirth: "01/01/01",
						Status:      customer.Active,
					}},
					nil,
				)
			},
			want: []dto.CustomerResponse{{
				Id:          "1",
				Name:        "Bob",
				City:        "Boston",
				Zipcode:     "90210",
				DateOfBirth: "01/01/01",
				Status:      "active",
			}},
			want1: nil,
		},
		{
			name: "Test got error",
			setUp: func(m *mock.MockcustomerData) {
				m.EXPECT().FindAll().Return(
					nil,
					errs.NewUnexpectedError("dummy err"))
			},
			want: nil,
			want1: &errs.AppError{
				Code:    http.StatusInternalServerError,
				Message: "dummy err",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			data := mock.NewMockcustomerData(ctrl)
			tt.setUp(data)

			s := CustomerService{
				data: data,
			}
			got, got1 := s.GetAllCustomers()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerService.GetAllCustomers() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CustomerService.GetAllCustomers() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCustomerService_GetCustomersByStatus(t *testing.T) {
	type args struct {
		status customer.Status
	}
	tests := []struct {
		name  string
		args  args
		setUp func(*mock.MockcustomerData, customer.Status)
		want  []dto.CustomerResponse
		want1 *errs.AppError
	}{
		{
			name: "Test get customer by status returns data",
			args: args{
				status: 0,
			},
			setUp: func(m *mock.MockcustomerData, s customer.Status) {
				m.EXPECT().ByActive(s).Return(
					[]customer.Customer{{
						Id:          "1",
						Name:        "Bob",
						City:        "Boston",
						Zipcode:     "90210",
						DateOfBirth: "01/01/01",
						Status:      customer.Active,
					}},
					nil,
				)
			},
			want: []dto.CustomerResponse{{
				Id:          "1",
				Name:        "Bob",
				City:        "Boston",
				Zipcode:     "90210",
				DateOfBirth: "01/01/01",
				Status:      "active",
			}},
			want1: nil,
		},
		{
			name: "Test got error",
			args: args{
				status: customer.Active,
			},
			setUp: func(m *mock.MockcustomerData, s customer.Status) {
				m.EXPECT().ByActive(s).Return(
					nil,
					errs.NewUnexpectedError("dummy err"),
				)
			},
			want: nil,
			want1: &errs.AppError{
				Code:    http.StatusInternalServerError,
				Message: "dummy err",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			data := mock.NewMockcustomerData(ctrl)
			tt.setUp(data, tt.args.status)

			s := CustomerService{
				data: data,
			}
			got, got1 := s.GetCustomersByStatus(tt.args.status)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerService.GetCustomersByStatus() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CustomerService.GetCustomersByStatus() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCustomerService_GetCustomer(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name  string
		args  args
		setUp func(*mock.MockcustomerData, string)
		want  *dto.CustomerResponse
		want1 *errs.AppError
	}{
		{
			name: "Test get customer by id returns data",
			args: args{
				id: "1",
			},
			setUp: func(m *mock.MockcustomerData, id string) {
				m.EXPECT().ById(id).Return(
					&customer.Customer{
						Id:          "1",
						Name:        "Bob",
						City:        "Boston",
						Zipcode:     "90210",
						DateOfBirth: "01/01/01",
						Status:      customer.Active,
					},
					nil,
				)
			},
			want: &dto.CustomerResponse{
				Id:          "1",
				Name:        "Bob",
				City:        "Boston",
				Zipcode:     "90210",
				DateOfBirth: "01/01/01",
				Status:      "active",
			},
			want1: nil,
		},
		{
			name: "Test got error",
			args: args{
				id: "1",
			},
			setUp: func(m *mock.MockcustomerData, id string) {
				m.EXPECT().ById(id).Return(
					nil,
					errs.NewUnexpectedError("dummy err"),
				)
			},
			want: nil,
			want1: &errs.AppError{
				Code:    http.StatusInternalServerError,
				Message: "dummy err",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			data := mock.NewMockcustomerData(ctrl)
			tt.setUp(data, tt.args.id)

			s := CustomerService{
				data: data,
			}
			got, got1 := s.GetCustomer(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerService.GetCustomer() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CustomerService.GetCustomer() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		cd customerData
	}
	tests := []struct {
		name string
		args args
		want CustomerService
	}{
		{
			name: "Test got new CustomerService",
			args: args{cd: mock.NewMockcustomerData(nil)},
			want: CustomerService{mock.NewMockcustomerData(nil)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.cd); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
