package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func GetLenguajesProgramacion(c *gin.Context) {
	dir, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener el directorio de trabajo"})
		return
	}
	filePath := filepath.Join(dir, "data", "categorias", "lenguajes-de-programacion.json")

	data, err := os.ReadFile(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var categoriaData interface{}
	if err := json.Unmarshal(data, &categoriaData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categoriaData)
}

func GetCategorias(c *gin.Context) {
	dir, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener el directorio de trabajo"})
		return
	}
	filePath := filepath.Join(dir, "data", "categorias", "categorias.json")

	data, err := os.ReadFile(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var categoriasData interface{}
	if err := json.Unmarshal(data, &categoriasData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categoriasData)
}
