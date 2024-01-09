package utils

import (
	"errors"
	"github.com/ktruedat/goBlockchain/internals/blockchain"
	"reflect"
)

const (
	nilStructFieldErr = "struct field cannot be nil"
)

func Validate[T blockchain.TransactionRequest](v T) error {
	if err := validateNilFields(v); err != nil {
		return err
	}
	return nil
}

// validateNilFields will go through the fields of a generic struct s
// and check whether it is nil or not
func validateNilFields[T blockchain.TransactionRequest](s T) error {
	sValue := reflect.ValueOf(s)
	sType := reflect.TypeOf(s)
	for i := 0; i < sType.NumField(); i++ {
		field := sValue.Field(i)

		// Check if the field is nil
		if field.Interface() == nil {
			return errors.New(nilStructFieldErr)
		}
	}
	return nil
}
