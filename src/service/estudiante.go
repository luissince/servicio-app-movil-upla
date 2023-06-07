package service

import (
	"context"
	"fmt"
	"servicio-app-movil-upla/src/database"
	"servicio-app-movil-upla/src/helper"
	"servicio-app-movil-upla/src/model"
)

var contx_estudiante = context.Background()

func ObtenerDatosEstudiante(codigo string) (model.DatosFicha, string) {
	datosFicha := model.DatosFicha{}

	// Crear conexión a la base de datos
	db, err := database.CreateConnection()
	if err != nil {
		return datosFicha, "No se puedo establecer una conexión."
	}
	defer db.Close()

	// Consulta SQL
	query := fmt.Sprintf("DECLARE @xml xml='<Dato><x_codigo>%s</x_codigo></Dato>'; EXEC [Ficha].[DatConAlu_S_1] @xml", codigo)
	// query := "DECLARE @xml xml='<Dato><x_codigo>" + codigo + "</x_codigo></Dato>'; EXEC [Ficha].[DatConAlu_S_1] @xml"

	// Ejecutar la consulta y escanear los resultados en el modelo
	row := db.QueryRow(query)
	if err := helper.ScanRow(row, &datosFicha); err != nil {
		return datosFicha, "No se puedo obtener los datos de la consulta"
	}

	return datosFicha, "ok"
}
