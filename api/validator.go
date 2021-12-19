package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/techschool/simplebank/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	// this syntax means taking the interface{} as string
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check if the currency is supported
		return util.IsSupportedCurrency(currency)
	}
	return false
}
