package hydraconfigurator

import (
	"errors"
	"reflect"
)

const (
	// CUSTOM is a handled but MarshallCustomConfig()
	CUSTOM uint8 = iota
)

var errWrongTypeError = errors.New("Type must be a pointer to a struct")

// GetConfiguration takes a constant representing the type, and empty interface and a filename
// it calls Marshall Custom Config to populste a struct with the contents of a file
func GetConfiguration(confType uint8, obj interface{}, filename string) (err error) {
	mysRValue := reflect.ValueOf(obj)
	if mysRValue.Kind() != reflect.Ptr || mysRValue.IsNil() {
		return errWrongTypeError
	}
	// get and confirm struct value
	mysRValue = mysRValue.Elem()
	// *object => object
	// reflection value of *object .Elem() => object() (Settable!!)
	if mysRValue.Kind() != reflect.Struct {
		return errWrongTypeError
	}

	switch confType {
	case CUSTOM:
		err = MarshallCustomConfig(mysRValue, filename)
	}
	return err
}
