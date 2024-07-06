package handler

import (
	"net/http"
	"os"
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

    router.ServeHTTP(w, r)
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }
    http.ListenAndServe(":" + port, http.HandlerFunc(Handler))
}
