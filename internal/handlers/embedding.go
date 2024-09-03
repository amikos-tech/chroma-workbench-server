package handlers

import "github.com/gin-gonic/gin"

func GetSupportedEmbeddingModels(c *gin.Context) {
	c.JSON(200, gin.H{"data": "Supported Embedding Models"})
}

type EmbeddingsRequest struct {
	Texts       []string               `json:"texts"`
	Vendor      string                 `json:"vendor"`
	Model       string                 `json:"model"`
	ModelConfig map[string]interface{} `json:"model_config"`
}

func EmbedText(c *gin.Context) {
	c.JSON(200, gin.H{"data": "Embed"})
}
