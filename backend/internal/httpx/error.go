package httpx

// import (
// 	"encoding/json"
// 	"net/http"
// 	// appErr "github.com/AsrofunNiam/lets-code-gantol-frozen-food-analyse/internal/errors"
// )

// type ErrorResponse struct {
// 	Code    int    `json:"code"`
// 	Message string `json:"message"`
// }

// func WriteError(w http.ResponseWriter, err error) {
// 	kind := appErr.UnwrapKind(err)

// 	status := appErr.HTTPCode(err)
// 	if status == 0 {
// 		status = http.StatusInternalServerError
// 	}

// 	msg := messageByKind(kind)

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(status)

// 	_ = json.NewEncoder(w).Encode(ErrorResponse{
// 		Code:    status,
// 		Message: msg,
// 	})
// }

// func messageByKind(k appErr.Kind) string {
// 	switch k {
// 	case appErr.KindBadRequest:
// 		return "invalid request"
// 	case appErr.KindNotFound:
// 		return "resource not found"
// 	case appErr.KindForbidden:
// 		return "forbidden"
// 	case appErr.KindTooManyRequest:
// 		return "too many requests"
// 	default:
// 		return "internal server error"
// 	}
// }
