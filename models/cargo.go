package models

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

type Cargo struct {
	Product  string  `json:"product"`
	SrcCity  string  `json:"origin_city"`
	SrcState string  `json:"origin_state"`
	SrcLat   float64 `json:"origin_lat"`
	SrcLng   float64 `json:"origin_lng"`
	DstCity  string  `json:"destination_city"`
	DstState string  `json:"destination_state"`
	DstLat   float64 `json:"destination_lat"`
	DstLng   float64 `json:"destination_lng"`
}

func ParseCargos(filename string) []Cargo {
	var cargos []Cargo
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
		//skip header record of csv
		if record[0] != "product" {
			cargos = createCargoFromRecord(cargos, record)
		}
	}

	return cargos
}

func createCargoFromRecord(cargos []Cargo, record []string) []Cargo {
	srcLat, _ := strconv.ParseFloat(record[3], 64)
	srcLng, _ := strconv.ParseFloat(record[4], 64)
	destLat, _ := strconv.ParseFloat(record[7], 64)
	destLng, _ := strconv.ParseFloat(record[8], 64)
	newCargo := Cargo{record[0], record[1], record[2], srcLat, srcLng, record[5], record[6], destLat, destLng}
	return append(cargos, newCargo)
}
