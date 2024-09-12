package core

import "errors"

var (
	ErrorNotAllowedForUserLevel = errors.New("not_allowed_for_user_level")
	ErrorValidation             = errors.New("validation_error")
)
