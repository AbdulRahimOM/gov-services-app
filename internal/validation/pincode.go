package validation

import (
	"strconv"

	"github.com/go-playground/validator/v10"
)

func init() {
	validate.RegisterValidation("pincode", validatePincode)
}

func validatePincode(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	if len(value) != 6 {
		return false
	}
	pinCodeNum, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return false
	}
	//check if decimal
	if pinCodeNum != float64(int64(pinCodeNum)) {
		return false
	}
	return (pinCodeNum >= 110000 && pinCodeNum <= 899999)
}
