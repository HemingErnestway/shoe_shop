package storage

import "reflect"

func updateFields[T any](current *T, new *T) {
	currentStruct := reflect.ValueOf(current).Elem()
	newStruct := reflect.ValueOf(new).Elem()

	for i := 0; i < newStruct.NumField(); i++ {
		newField := newStruct.Type().Field(i)
		if !newStruct.FieldByName(newField.Name).IsZero() {
			currentField := currentStruct.FieldByName(newField.Name)
			currentField.Set(newStruct.FieldByName(newField.Name))
		}
	}
}
