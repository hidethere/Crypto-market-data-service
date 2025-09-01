package response

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrSymbolNotFound = errors.New("symbol not found")
)

func Success(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(APIResponse{
		Success: true,
		Data:    data,
	})
}

func Error(w http.ResponseWriter, err error) {

	var msg string
	var status int

	switch {
	case errors.Is(err, ErrSymbolNotFound):
		msg = ErrSymbolNotFound.Error()
		status = http.StatusNotFound
	default:
		msg = "Internal server error"
		status = http.StatusInternalServerError
	}

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(APIResponse{Success: false, Error: msg})

}
