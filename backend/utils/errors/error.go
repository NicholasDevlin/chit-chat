package errors

import "errors"

var (
	ERR_CREATE_CUSTOMER  = errors.New("Failed to save customer")
	ERR_CREATE_PRODUCT   = errors.New("Failed to save product")
	ERR_ADDRESS_IS_EMPTY = errors.New("Address is empty")
	ERR_NAME_IS_EMPTY    = errors.New("Name is empty")
	ERR_AGE_IS_EMPTY     = errors.New("Age is empty")
	ERR_USER_NOT_FOUND   = errors.New("User not found")
	ERR_DELETE_DATA = errors.New("Failed felete data")
	ERR_EMAIL_IS_EMPTY = errors.New("Email is empty")
	ERR_WRONG_PASSWORD = errors.New("Wrong password")
	ERR_PASSWORD_IS_EMPTY = errors.New("Password is empty")
	ERR_PHONE_NUMBER_IS_EMPTY = errors.New("Phone number is empty")
	ERR_EMAIL_IS_TAKEN = errors.New("Email is taken")
	ERR_BCRYPT_PASSWORD = errors.New("Failed to login")
	ERR_SAVE_DATA = errors.New("Save data failed")
	ERR_TOKEN = errors.New("Failed to login")
	ERR_UNAUTHORIZE = errors.New("Unauthorized user")
)
