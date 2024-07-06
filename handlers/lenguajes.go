package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetLenguaje(c *gin.Context) {
    lenguaje := c.Param("lenguaje")
    data, err := os.ReadFile("./data/lenguajes-de-programacion/" + lenguaje + ".json")
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
