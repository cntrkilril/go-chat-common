package errors

import "errors"

type (
	Error struct {
		msg  string
		code ErrCode
	}
	ErrCode int64
)

const (
	_ = iota
	ErrCodeUnknown
	ErrCodeNotFound
	ErrCodeInvalidArgument
	ErrCodeAlreadyExist
)

func (e *Error) Error() string {
	return e.msg
}

func (e *Error) Code() ErrCode {
	return e.code
}

var _ error = &Error{}

func NewError(msg string, code ErrCode) *Error {
	return &Error{msg, code}
}

var (
	ErrUnknown             = errors.New("что-то пошло не так")
	ErrValidationError     = errors.New("невалидные данные")
	ErrUserNotFound        = errors.New("пользователь не найден")
	ErrUserAlreadyExist    = errors.New("пользователь уже существует")
	ErrSessionNotFound     = errors.New("сессия не найдена")
	ErrSessionAlreadyExist = errors.New("сессия уже существует")
)
