package app

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/bradrogan/banking/errs"
)

func Test_writeResponse(t *testing.T) {
	type args struct {
		w    http.ResponseWriter
		code int
		data any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writeResponse(tt.args.w, tt.args.code, tt.args.data)
		})
	}
}

func Test_readRequest(t *testing.T) {
	type args struct {
		r   *http.Request
		dto any
	}
	tests := []struct {
		name string
		args args
		want *errs.AppError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readRequest(tt.args.r, tt.args.dto); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
