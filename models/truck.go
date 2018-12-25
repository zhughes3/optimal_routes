package models

type Truck struct {
	Truck string `json:"truck"`
	City string `json:"city"`
	State string `json:"state"`
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
