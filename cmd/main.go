package cmd

import (
	"chroma-workbench-server/internal/handlers"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} //TODO make this configurable
	router.Use(cors.New(config))

	router.GET("/embedding/models", handlers.GetSupportedEmbeddingModels)
	router.POST("/embedding/text/embed", handlers.EmbedText)
	err := router.Run(":8089")
	if err != nil {
		fmt.Printf("Error starting server: %s", err.Error())
		return
	}
}
