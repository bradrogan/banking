package app

import (
	"encoding/json"
	"net/http"

	"github.com/bradrogan/banking/errs"
	"github.com/bradrogan/banking/logger"
	"go.uber.org/zap"
)

func writeResponse(w http.ResponseWriter, code int, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}

}

func readRequest(r *http.Request, dto any) *errs.AppError {
	err := json.NewDecoder(r.Body).Decode(dto)
	if err != nil {
		logger.Error("error reading request body", zap.Error(err))
		return errs.NewUnexpectedError("error reqding request body")
	}
	return nil
}

