package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"log"
	"strconv"
	"routes/models"
)

func main() {
	trucks := readTrucks("trucks.csv")
	cargo := readCargos("cargo.csv")

	fmt.Println("Trucks: ")
	fmt.Println(trucks)

	fmt.Println("Cargo: ")
	fmt.Println(cargo)
}

func readTrucks(filename string) []models.Truck {
	var trucks []models.Truck
	f, _ := os.Open(filename)
	defer f.Close()
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if record[0] != "truck" {
			trucks = createTruckFromRecord(trucks, record)
		}
	}

	return trucks
}

func createTruckFromRecord(trucks []models.Truck, record []string) []models.Truck {
	lat, _ := strconv.ParseFloat(record[3], 64)
	lng, _ := strconv.ParseFloat(record[4], 64)
	newTruck := models.Truck{record[0], record[1], record[2], lat, lng}
	return append(trucks, newTruck)
}

func readCargos(filename string) []models.Cargo {
	var cargos []models.Cargo
	f, _ := os.Open(filename)
	defer f.Close()
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if record[0] != "product" {
			cargos = createCargoFromRecord(cargos, record)
		}
	}

	return cargos
}

func createCargoFromRecord(cargos []models.Cargo, record []string) []models.Cargo {
	srcLat, _ := strconv.ParseFloat(record[3], 64)
	srcLng, _ := strconv.ParseFloat(record[4], 64)
	destLat, _ := strconv.ParseFloat(record[7], 64)
	destLng, _ := strconv.ParseFloat(record[8], 64)
	newCargo := models.Cargo{record[0], record[1], record[2], srcLat, srcLng, record[5], record[6], destLat, destLng}
	return append(cargos, newCargo)
}