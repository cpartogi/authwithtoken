package utils

import "authwithtoken/lib/helper"

func errorType(err error) (int, error) {
	return helper.CommonError(err)
}
