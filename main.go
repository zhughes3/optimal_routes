package main

import (
	"fmt"
	"github.com/zhughes3/optimal_routes/models"
)

func main() {
	trucks := models.ParseTrucks("trucks.csv")
	cargo := models.ParseCargos("cargo.csv")

	fmt.Println("Trucks: ")
	fmt.Println(trucks)

	fmt.Println("Cargo: ")
	fmt.Println(cargo)
}
