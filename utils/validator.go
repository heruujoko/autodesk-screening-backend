package utils

import (
	"fmt"
	"sync"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate
var onceValidator sync.Once

func init() {
	onceValidator.Do(func() {
		validate = validator.New()
	})
}

type FieldError struct {
	FieldName 	string `json:"field_name"`
	Status 		string `json:"status"`
	Type 		string `json:"type"`
}

type ValidationResponse struct {
	Errors []FieldError	`json:"errors"`
}

func ValidateSchema(data interface{}) (error, interface{}) {
	err := validate.Struct(data)
	if (err != nil) {
		var errs = []FieldError{}
		formError := ValidationResponse{
			Errors: errs,
		}

		for _, err := range err.(validator.ValidationErrors) {

			/** 
				more fileds to explore in :
			 	https://github.com/go-playground/validator/blob/master/_examples/simple/main.go
			**/

			fld := FieldError {
				FieldName: err.Field(),
				Status: err.Tag(),
				Type: fmt.Sprintf("%v", err.Kind()),
			}
			formError.Errors = append(formError.Errors, fld)
		}

		return err, formError
	}

	return nil, nil
}