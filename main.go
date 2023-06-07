package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	time.LoadLocation("America/Lima")
	godotenv.Load()

	var go_port string = os.Getenv("GO_PORT")

	var ruta_log string = os.Getenv("RUTA_LOG")

	// Crear archivo log
	f, err := os.Create(ruta_log)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	router := gin.Default()
	router.Use(gin.Logger())

	// Middleware para CORS
	router.Use(cors.Default())

	// Rutas
	router.GET("/", func(c *gin.Context) {
		log.Println("Endpoint ping")
		c.JSON(http.StatusOK, gin.H{
			"message": "API GO LANG APP MOVIL UPLA",
		})
	})

	router.Run(go_port)
}
