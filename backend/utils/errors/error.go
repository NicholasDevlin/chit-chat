package errors

import "errors"

var (
	ERR_CREATE_CUSTOMER = errors.New("Failed to save customer")
	ERR_ADDRESS_IS_EMPTY = errors.New("Address is empty")
	ERR_NAME_IS_EMPTY = errors.New("Name is empty")
	ERR_AGE_IS_EMPTY = errors.New("Age is empty")
)
