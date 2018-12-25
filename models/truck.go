package models

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

type Truck struct {
	Truck string  `json:"truck"`
	City  string  `json:"city"`
	State string  `json:"state"`
	Lat   float64 `json:"lat"`
	Lng   float64 `json:"lng"`
}

func ParseTrucks(filename string) []Truck {
	var trucks []Truck
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
		// skip header record of csv
		if record[0] != "truck" {
			trucks = createTruckFromRecord(trucks, record)
		}
	}

	return trucks
}

func createTruckFromRecord(trucks []Truck, record []string) []Truck {
	lat, _ := strconv.ParseFloat(record[3], 64)
	lng, _ := strconv.ParseFloat(record[4], 64)
	newTruck := Truck{record[0], record[1], record[2], lat, lng}
	return append(trucks, newTruck)
}
