package handler

import (
	"net/http"
	"os"
	"path/filepath"
	"recursos/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    gin.SetMode(gin.ReleaseMode)
    router := gin.New()
    router.Use(gin.Recovery())
    router.Use(cors.New(cors.Config{
        AllowAllOrigins:  true,
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    router.GET("/api/categorias/lenguajes-de-programacion", handlers.GetLenguajesProgramacion)
    router.GET("/api/categorias/lenguajes-de-programacion/:lenguaje", handlers.GetLenguaje)
    router.GET("/api/categorias", handlers.GetCategorias)

    // Endpoint temporal para verificar archivos
    router.GET("/api/check-files", func(c *gin.Context) {
        dir, err := os.Getwd()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener el directorio de trabajo"})
            return
        }
        dataDir := filepath.Join(dir, "data")
        files, err := os.ReadDir(dataDir)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        fileNames := []string{}
        for _, file := range files {
            fileNames = append(fileNames, file.Name())
        }

        c.JSON(http.StatusOK, gin.H{"files": fileNames})
    })

    router.ServeHTTP(w, r)
}
