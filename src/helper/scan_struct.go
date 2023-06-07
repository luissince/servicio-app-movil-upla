package helper

import (
	"database/sql"
	"reflect"

	_ "github.com/microsoft/go-mssqldb"
)

func ScanRow(row *sql.Row, model interface{}) error {
	values := make([]interface{}, 0)
	modelValue := reflect.ValueOf(model).Elem()
	modelType := modelValue.Type()

	// Recorrer los campos del modelo y crear una lista de punteros a los campos
	for i := 0; i < modelType.NumField(); i++ {
		fieldValue := modelValue.Field(i)
		values = append(values, fieldValue.Addr().Interface())
	}

	// Escanear los valores de las columnas en los campos del modelo
	err := row.Scan(values...)
	if err != nil {
		return err
	}

	return nil
}

func ScanRows(row *sql.Rows, dest interface{}) error {
	columns, err := row.Columns()
	if err != nil {
		return err
	}

	values := make([]interface{}, len(columns))
	columnPointers := make([]interface{}, len(columns))

	for i := range columns {
		columnPointers[i] = &values[i]
	}

	err = row.Scan(columnPointers...)
	if err != nil {
		return err
	}

	destValue := reflect.ValueOf(dest).Elem()
	for i, columnValue := range values {
		fieldValue := destValue.FieldByName(columns[i])
		if fieldValue.IsValid() {
			field := fieldValue.Addr().Interface()
			if scanner, ok := field.(sql.Scanner); ok {
				err := scanner.Scan(columnValue)
				if err != nil {
					return err
				}
			} else {
				destValue.FieldByName(columns[i]).Set(reflect.ValueOf(columnValue))
			}
		}
	}

	return nil
}
