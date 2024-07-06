package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func GetLenguaje(c *gin.Context) {
    lenguaje := c.Param("lenguaje")
    dir, err := os.Getwd()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener el directorio de trabajo"})
        return
    }
    filePath := filepath.Join(dir, "data", "lenguajes-de-programacion", lenguaje+".json")

    data, err := os.ReadFile(filePath)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    var lenguajeData interface{}
    if err := json.Unmarshal(data, &lenguajeData); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, lenguajeData)
}
