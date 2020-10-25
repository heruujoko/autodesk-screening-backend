package utils

import (
	"github.com/labstack/echo/v4"
)

func ParseAndValidate(context echo.Context, object interface{}) *ErrorBag {
	errors := ErrorBag{}
	// parse
	if err := context.Bind(object); err != nil {
		 errors.Code = "BAD_REQUEST"
		 errors.Message = err.Error()
		 return &errors
	}

	// validate
	validator, errObj := ValidateSchema(object)
	if (validator != nil) {
		errors.Code = "BAD_REQUEST"
		errors.Message = validator.Error()
		errors.Data = errObj
		return &errors
	}

	return nil
}
