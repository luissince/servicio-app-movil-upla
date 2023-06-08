package service

import (
	"database/sql"
	"fmt"
	"servicio-app-movil-upla/src/database"
	"servicio-app-movil-upla/src/helper"
	"servicio-app-movil-upla/src/model"
)

func ObtenerDatosEstudiante(codigo string) (model.DatosFicha, string) {
	datosFicha := model.DatosFicha{}

	// Crear conexión a la base de datos
	db, err := database.CreateConnection()
	if err != nil {
		return datosFicha, "No se puedo establecer una conexión."
	}
	defer db.Close()

	// Consulta SQL
	declarar := "DECLARE @xml xml='<Dato><x_codigo>%s</x_codigo></Dato>';"
	procedimiento := "EXEC [Ficha].[DatConAlu_S_1] @xml"
	query := fmt.Sprintf(declarar+procedimiento, codigo)
	// query := "DECLARE @xml xml='<Dato><x_codigo>" + codigo + "</x_codigo></Dato>'; EXEC [Ficha].[DatConAlu_S_1] @xml"

	// Ejecutar la consulta y escanear los resultados en el modelo
	row := db.QueryRow(query)
	if err := helper.ScanRow(row, &datosFicha); err != nil {
		return datosFicha, "No se puedo obtener los datos de la consulta"
	}

	return datosFicha, "ok"
}

func ObtenerProgresoAcademico(codigo string) ([]model.ProgresoAcademico, string) {
	progresoAcademicos := []model.ProgresoAcademico{}

	// Crear conexión a la base de datos
	db, err := database.CreateConnection()
	if err != nil {
		return progresoAcademicos, "No se puedo establecer una conexión."
	}
	defer db.Close()

	query := "exec [Alumno].[ObtenerDniPorCodigo] @codigo"
	row := db.QueryRow(query, sql.Named("codigo", codigo))
	var dni string
	err = row.Scan(&dni)
	if err == sql.ErrNoRows || err != nil {
		fmt.Println(err)
		return progresoAcademicos, "El código ingresado no tiene ningún dni asociado."
	}

	query = "exec [Alumno].ObtenerPregradoPorDni @dni"
	rows, err := db.Query(query, sql.Named("dni", dni))

	defer rows.Close()

	for rows.Next() {
		progresoAcademico := model.ProgresoAcademico{}
		rows.Scan(&progresoAcademico.Codigo, &progresoAcademico.Descripcion)
		progresoAcademicos = append(progresoAcademicos, progresoAcademico)
	}

	return progresoAcademicos, "ok"
}
