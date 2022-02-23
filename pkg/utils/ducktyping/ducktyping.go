package ducktyping

import (
	"fmt"
	"reflect"
)

func DuckCopy(from interface{}, to interface{}) error {
	fromR := reflect.ValueOf(from)
	if fromR.Kind() == reflect.Ptr {
		fromR = fromR.Elem()
	}
	toR := reflect.ValueOf(to)

	if toR.Kind() != reflect.Ptr {
		return fmt.Errorf("DuckCopy: to parameter must be a pointer")
	}

	for i := 0; i < fromR.NumField(); i++ {
		fromty := fromR.Type().Field(i)
		fromfi := fromR.Field(i)
		tofi := toR.Elem().FieldByName(fromty.Name)
		if !tofi.IsValid() {
			return fmt.Errorf("DuckCopy: field %s does not exist in 'to'", fromty.Name)
		}
		tofi.Set(fromfi)
	}
	return nil
}

func AddFieldToInterface(data interface{}, fieldsToAdd []reflect.StructField) interface{} {
	dataR := reflect.ValueOf(data)
	if dataR.Kind() == reflect.Ptr {
		dataR = dataR.Elem()
	}

	var fields []reflect.StructField
	for i := 0; i < dataR.NumField(); i++ {
		dataty := dataR.Type().Field(i)
		fields = append(fields,
			reflect.StructField{
				Name: dataty.Name,
				Type: dataty.Type,
				Tag:  dataty.Tag,
			})
	}
	fields = append(fields, fieldsToAdd...)
	newty := reflect.StructOf(fields)
	return reflect.New(newty).Interface()
}

func IsFieldInInterface(data interface{}, field string) bool {
	dataR := reflect.ValueOf(data)
	if dataR.Kind() == reflect.Ptr {
		return dataR.Elem().FieldByName(field).IsValid()
	}
	return dataR.FieldByName(field).IsValid()
}

func AffectValue(data interface{}, field string, value interface{}) {
	reflect.ValueOf(data).Elem().FieldByName(field).Set(reflect.ValueOf(value))
}
