package helper

import (
	"authwithtoken/lib/constant"
	"fmt"
	"net/http"
)

var commonErrorMap = map[error]int{
	constant.ErrNotFound:    http.StatusNotFound,
	constant.ErrConflict:    http.StatusConflict,
	constant.ErrBadRequest:  http.StatusBadRequest,
	constant.ErrInvalidUuid: http.StatusBadRequest,
}

// CommonError is
func CommonError(err error) (int, error) {
	switch err {
	case constant.ErrNotFound:
		return commonErrorMap[constant.ErrNotFound], constant.ErrNotFound
	case constant.ErrConflict:
		return commonErrorMap[constant.ErrConflict], constant.ErrConflict
	case constant.ErrBadRequest:
		return commonErrorMap[constant.ErrBadRequest], constant.ErrBadRequest
	case constant.ErrInvalidUuid:
		return commonErrorMap[constant.ErrBadRequest], constant.ErrInvalidUuid
	case constant.ErrTitle:
		return commonErrorMap[constant.ErrBadRequest], constant.ErrTitle
	case constant.ErrTypeNotFound:
		return commonErrorMap[constant.ErrBadRequest], constant.ErrTypeNotFound
	}
	return http.StatusInternalServerError, fmt.Errorf(err.Error())
}
