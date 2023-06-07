package controller

import (
	"net/http"
	"servicio-app-movil-upla/src/model"
	"servicio-app-movil-upla/src/service"

	"github.com/gin-gonic/gin"
)

func ObtenerDatosEstudiante(c *gin.Context) {
	// Obtener el idConsulta de la URL
	codigo := c.Param("codigo")
	if codigo == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "No se pudo procesar el parametro de la URL."})
		return
	}

	datosFicha, rpta := service.ObtenerDatosEstudiante(codigo)
	if rpta != "ok" {
		c.IndentedJSON(http.StatusBadRequest, model.Error{Message: rpta})
		return
	}

	c.IndentedJSON(http.StatusOK, datosFicha)
}
