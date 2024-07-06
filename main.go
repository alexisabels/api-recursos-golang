package main

import (
	"os"
	"recursos/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.Use(cors.New(cors.Config{
        AllowAllOrigins:  true,
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    r.GET("/api/categorias/lenguajes-de-programacion", handlers.GetLenguajesProgramacion)
    r.GET("/api/categorias/lenguajes-de-programacion/:lenguaje", handlers.GetLenguaje)
	r.GET("/api/categorias", handlers.GetCategorias)

    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }
    r.Run(":" + port)
}
