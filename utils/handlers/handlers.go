package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/yesseneon/bookstore-utils/errors"
)

func RespondJson(w http.ResponseWriter, code int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(body)
}

func RespondError(w http.ResponseWriter, restErr *errors.RESTError) {
	RespondJson(w, restErr.Status, restErr)
}
