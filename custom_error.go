package main

type ErrCode string

func (code ErrCode) Wrap(err error, message string) error {
	return &CustomError{ErrCode: code, Message: message, Err: err}
}

const (
	// R00x is Repository errors
	DuplicateKeyErr ErrCode = "R001" // SQL error
)

type CustomError struct {
	ErrCode        // error code for showing response and logging
	Message string // error message for showing response
	Err     error  // internal error for error chain
}

func (ce *CustomError) Error() string {
	return ce.Message
}

// for using errors.Is/errors.As
func (ce *CustomError) Unwrap() error {
	return ce.Err
}
