package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetLenguajesProgramacion(c *gin.Context) {
    data, err := os.ReadFile("./data/categorias/lenguajes-de-programacion.json")
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
    data, err := os.ReadFile("./data/categorias/categorias.json") 
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

