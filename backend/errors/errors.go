package apperr

import (
	"errors"
	"fmt"
	"net/http"
)

type errorCode string

// TODO: константы это не тип errorCode
// Коды ошибок
const (
	ErrNotFound     = "not_found"
	ErrInternal     = "internal"
	ErrUnauthorized = "unauthorized"
	ErrConflict     = "conflict"
	ErrValidation   = "validation_error"
	ErrExpired      = "expired"
)

type Error struct {
	// Вложенная ошибка
	Err error `json:"error"`
	// Дополнительный контекст (трассировка)
	Fields map[string]interface{}
	// Код ошибки
	Code errorCode `json:"code"`
	// Сообщение, понятное пользователю
	Message string `json:"message"`
	// Выполняемая операция
	Op string `json:"op"`
}

func New(err error, code errorCode, msg string, op string) (e *Error) {
	return &Error{
		Err:     err,
		Code:    code,
		Message: msg,
		Fields:  make(map[string]interface{}),
		Op:      op,
	}
}

func (e Error) Error() string {
	return fmt.Sprintf("message: %s, code: %s, op: %s error: %s", e.Message, e.Code, e.Op, e.Err)
}

func ErrorCode(err error) errorCode {
	if error, ok := err.(*Error); ok {
		return error.Code
	}
	return ErrInternal
}

// Соответствие кода ошибки и HTTP-статуса
var codeToHTTPStatusMap = map[errorCode]int{
	ErrNotFound:     http.StatusNotFound,
	ErrInternal:     http.StatusInternalServerError,
	ErrUnauthorized: http.StatusUnauthorized,
	ErrConflict:     http.StatusConflict,
	ErrValidation:   http.StatusBadRequest,
	ErrExpired:      http.StatusUnauthorized,
}

func ErrCodeToHTTPStatus(err error) int {
	code := ErrorCode(err)
	if v, ok := codeToHTTPStatusMap[code]; ok {
		return v
	}

	// Возвращаем стандартный HTTP-статус для неизвестных ошибок
	return http.StatusInternalServerError
}

func Message(err error) string {
	e := &Error{}
	if errors.As(err, e) {
		return e.Message
	}
	return err.Error()
}

func IsErrorCode(err error, code errorCode) bool {
	errcode := ErrorCode(err)
	return errcode == code
}
