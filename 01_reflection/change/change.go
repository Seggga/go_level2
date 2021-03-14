package change

import (
	"fmt"
	"reflect"
)

func SetValues(in interface{}, values map[string]interface{}) error {
	// check passed arguments
	if in == nil {
		return fmt.Errorf("passed nil (argument in), while a pointer to struct-type was expected")
	}
	val := reflect.ValueOf(in)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("passed non-structure-type")
	}

	structType := val.Type()
	for i := 0; i < val.NumField(); i++ {
		newValue, ok := values[structType.Field(i).Name]
		if !ok {
			// there is no new value for this field in the structure
			continue
		}
		// compare types
		currentField := val.Field(i)
		fieldType := currentField.Type()
		mapDataType := reflect.TypeOf(newValue)
		if fieldType != mapDataType {
			return fmt.Errorf("value's type mismatch in %s : %s expected, %s was passed", structType.Field(i).Name, fieldType, mapDataType)
		}

		// change the value of current struct field
		switch fieldType.Kind() {
		case reflect.Int64:
			currentField.SetInt(newValue.(int64))
		case reflect.Uint64:
			currentField.SetUint(newValue.(uint64))
		case reflect.Float64:
			currentField.SetFloat(newValue.(float64))
		case reflect.String:
			currentField.SetString(newValue.(string))
		case reflect.Bool:
			currentField.SetBool(newValue.(bool))
		default:
			return fmt.Errorf("Unsupported value's type (%s was passed)", mapDataType)
		}

	}

	return nil
}
