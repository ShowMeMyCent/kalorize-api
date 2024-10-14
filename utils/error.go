package utils

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

const (
	ErrGeneral  = 500
	ErrDatabase = 504
)

// ApplicationError is a special structure for handling errors within the application
type ApplicationError struct {
	msg   string // error message
	code  string // error code
	cause error  // the original error that caused the ApplicationError
}

// getErrorMessage returns the error message
func (e *ApplicationError) getErrorMessage() string {
	return e.msg
}

// error returns the error message
func (e *ApplicationError) Error() string {
	msg := e.getErrorMessage()
	if e.code != "" {
		msg = fmt.Sprintf("%s (type: %s)", msg, e.code)
	}
	if e.cause != nil {
		msg = fmt.Sprintf("%s: %v", msg, e.cause)
	}
	return msg
}

// getErrorCode returns the error code
func (e *ApplicationError) getErrorCode() string {
	return e.code
}

// NewApplicationError creates a new error with a message and code, with an optional cause
func NewApplicationError(msg string, code string, cause error) error {
	applicationErr := &ApplicationError{
		msg:   msg,
		code:  code,
		cause: cause,
	}

	return applicationErr
}

// SetApplicationError creates a new error with a message and code, with an optional cause
func SetApplicationError(errMessage, errCode string, cause ...error) error {
	var errCause error
	if len(cause) > 0 {
		errCause = cause[0]
	}
	return NewApplicationError(errMessage, errCode, errCause)
}

// SetApplicationErrorWithCause creates a new error with a message, code, and cause
func SetApplicationErrorWithCause(errMessage, errCode string, cause error) error {
	return NewApplicationError(errMessage, errCode, cause)
}

// Error creates a new error with additional information about the file and line where the error occurred
func Error(err error, errCode int, errMessage string) error {

	if _, file, line, ok := runtime.Caller(1); ok {
		f := strings.Split(file, "/")
		errMessage = fmt.Sprintf("%s/%s:%d", errMessage, f[len(f)-1], line)
	}
	if err == nil {
		return SetApplicationError(errMessage, ErrorCodeString(errCode))
	}
	return SetApplicationErrorWithCause(errMessage, ErrorCodeString(errCode), err)
}

// ErrorCodeString converts an integer error code into a string
func ErrorCodeString(errCode int) string {
	return fmt.Sprintf("%d", errCode)
}

// ErrorMessage returns the error message from the given error
func ErrorMessage(err error) string {
	return err.Error()
}

// ErrorCode returns the error code from the given error
func ErrorCode(err error) int {

	var applicationErr *ApplicationError

	// Check if the error can be asserted to an ApplicationError
	if errors.As(err, &applicationErr) {
		code, _ := strconv.Atoi(applicationErr.getErrorCode())
		return code

	}

	return ErrGeneral
}

// ErrorCodeAndMessage returns the error code and error message from the given error
func ErrorCodeAndMessage(err error) (code int, remark string) {
	return ErrorCode(err), ErrorMessage(err)
}
