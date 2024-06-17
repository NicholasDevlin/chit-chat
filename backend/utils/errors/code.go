package errors

import "net/http"


func GetCodeError(err error) int {
	switch err {
	case ERR_ADDRESS_IS_EMPTY:
		return http.StatusBadRequest
	case ERR_AGE_IS_EMPTY:
		return http.StatusBadRequest
	case ERR_NAME_IS_EMPTY:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
