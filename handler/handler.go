package handler

import "github.com/gin-gonic/gin"

func CreateopeningHandler(ctz *gin.Context) {
	ctz.JSON(200, gin.H{
		"message": "Create opening",
	})
}

func UpdateopeningHandler(ctz *gin.Context) {
	ctz.JSON(200, gin.H{
		"message": "Update opening",
	})
}

func DeleteopeningHandler(ctz *gin.Context) {
	ctz.JSON(200, gin.H{
		"message": "Delete opening",
	})
}
func ShowopeningHandler(ctz *gin.Context) {
	ctz.JSON(200, gin.H{
		"message": "Show opening",
	})
}

func ShowallopeningHandler(ctz *gin.Context) {
	ctz.JSON(200, gin.H{
		"message": "Show all openings",
	})
}
