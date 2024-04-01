package constant

import "fmt"

const (
	PhoneNumberCharLength = "phoneNumber must be at minimum 10 characters and maximum 13 characters"
	PhoneNumberIndonesian = "phoneNumber must start with the Indonesia country code +62"
	PassWordCharLength    = "password must be minimum 6 characters and maximum 64 characters"
	FullNameCharLength    = "fullName must be at minimum 3 characters and maximum 60 characters"
	PasswordReqChar       = "password containing at least 1 capital characters AND 1 number AND 1 special (nonalpha-numeric) character"
	InternalServerError   = "Internal Server Error"
	PhoneNumberRequired   = "phoneNumber is required"
	PasswordRequired      = "password is required"
	DataNotFound          = "data not found"
	InvalidPassword       = "invalid password"
	DuplicatePhone        = "phone number already exist"
)

var (
	ErrNotFound     = fmt.Errorf("data not found")
	ErrConflict     = fmt.Errorf("conflict, data already exist")
	ErrBadRequest   = fmt.Errorf("bad request")
	ErrTitle        = fmt.Errorf("title required")
	ErrTypeNotFound = fmt.Errorf("tutorial Type not found")
)
