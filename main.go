package main

import (
	"github.com/zhughes3/optimal_routes/models"
)

func main() {
	var matrix = models.InitMatrix()
	minRoute := models.FindOptimalRoutes(matrix)
	minRoute.Print()
}