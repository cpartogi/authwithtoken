package helper

import (
	"authwithtoken/lib/constant"
	"fmt"
	"net/http"
	"strings"
)

var commonErrorMap = map[error]int{
	constant.ErrNotFound:   http.StatusNotFound,
	constant.ErrConflict:   http.StatusConflict,
	constant.ErrBadRequest: http.StatusBadRequest,
}

// CommonError is
func CommonError(err error) (int, error) {

	if strings.Contains(err.Error(), "required") || strings.Contains(err.Error(), "character") || strings.Contains(err.Error(), "email") || strings.Contains(err.Error(), "password") {
		return http.StatusBadRequest, fmt.Errorf(err.Error())
	}

	switch err {
	case constant.ErrNotFound:
		return commonErrorMap[constant.ErrNotFound], constant.ErrNotFound
	case constant.ErrConflict:
		return commonErrorMap[constant.ErrConflict], constant.ErrConflict
	case constant.ErrBadRequest:
		return commonErrorMap[constant.ErrBadRequest], constant.ErrBadRequest
	}
	return http.StatusInternalServerError, fmt.Errorf(err.Error())
}
