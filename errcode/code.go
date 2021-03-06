package errcode

import "net/http"

var (
	// OK success response
	OK = &ErrCode{
		Status:  http.StatusOK,
		Code:    "OK",
		Message: "OK",
	}

	// Created record created
	Created = &ErrCode{
		Status:  http.StatusCreated,
		Code:    "OK",
		Message: "created",
	}

	// follows are error responses of server-side

	// InternalServerError unknown error inside server
	InternalServerError = &ErrCode{
		Status:  http.StatusInternalServerError,
		Code:    "InternalError",
		Message: "Internal server error.",
	}

	// CustomInternalServerError custom error message of response
	// Usage:
	// err := errcode.New(errcode.CustomInternalServerError, nil)
	// err1 := err.Add("some error message string")
	// handler.SendResponse(c, err1, nil)
	CustomInternalServerError = &ErrCode{
		Status:  http.StatusInternalServerError,
		Code:    "InternalError",
		Message: "Internal server error: %v",
	}

	// ServerPanicError panic error
	ServerPanicError = &ErrCode{
		Status:  http.StatusInternalServerError,
		Code:    "PanicError",
		Message: "Server meets up panic error, please contact admin.",
	}

	// follows are error responses of client-side

	// ValidationError params, queries, form datas validation
	ValidationError = &ErrCode{
		Status:  http.StatusBadRequest,
		Code:    "ValidationError",
		Message: "%v",
	}

	// UserNotFoundError user not found
	UserNotFoundError = &ErrCode{
		Status:  http.StatusNotFound,
		Code:    "UserNotFound",
		Message: "The user is not found.",
	}

	// PasswordIncorrectError password of user is wrong
	PasswordIncorrectError = &ErrCode{
		Status:  http.StatusBadRequest,
		Code:    "PasswordIncorrect",
		Message: "The password is incorrect.",
	}

	// TokenInvalidError The token carried by the request is invalid.
	TokenInvalidError = &ErrCode{
		Status:  http.StatusUnauthorized,
		Code:    "TokenInvalid",
		Message: "%v",
	}

	// APISignError The signature carried by the request is invalid.
	APISignError = &ErrCode{
		Status:  http.StatusUnauthorized,
		Code:    "SignatureInvalid",
		Message: "%v",
	}

	// AccessForbiddenError The client ip or the server api is disallowed to access.
	AccessForbiddenError = &ErrCode{
		Status:  http.StatusForbidden,
		Code:    "AccessForbidden",
		Message: "%s",
	}

	// RecordNotFoundError The data of user request not exists.
	RecordNotFoundError = &ErrCode{
		Status:  http.StatusNotFound,
		Code:    "RecordNotFound",
		Message: "Record not found.",
	}

	// TooManyRequestError The client ip access frequency exceeds limit.
	TooManyRequestError = &ErrCode{
		Status:  http.StatusTooManyRequests,
		Code:    "TooManyRequest",
		Message: "Too many request.",
	}
)
