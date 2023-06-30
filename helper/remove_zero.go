package helper

import (
	"reflect"
)

func RemoveZero(data interface{}, dest *map[string]interface{}) {
	value := reflect.ValueOf(data).Elem()
	typeOfData := value.Type()

	for i := 0; i < value.NumField(); i++ {
		fieldValue := value.Field(i)
		fieldType := typeOfData.Field(i)
		zeroValue := reflect.Zero(fieldValue.Type())

		if !reflect.DeepEqual(fieldValue.Interface(), zeroValue.Interface()) {
			fieldName := fieldType.Tag.Get("bson")
			(*dest)[fieldName] = fieldValue.Interface()
		}

	}
}
