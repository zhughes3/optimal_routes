package main

import (
	"fmt"
	"github.com/zhughes3/optimal_routes/distance"
	"github.com/zhughes3/optimal_routes/models"
	"math"
)

type Distance struct {
	truck    string
	product  string
	distance float64
	isUsed   bool
}

func main() {
	var matrix = initMatrix()
	minRoute := findOptimalRoute(matrix)
	if minRoute != nil {
		fmt.Println(minRoute)
	} else {
		fmt.Println("error")
	}
}

func initMatrix() map[string][]Distance {
	trucks := models.ParseTrucks("trucks.csv")
	cargo := models.ParseCargos("cargo.csv")

	routes := make(map[string][]Distance)
	for _, truck := range trucks {
		var truckDistances []Distance
		var lat, lng = truck.Lat, truck.Lng
		for _, c := range cargo {
			var clat, clng = c.SrcLat, c.SrcLng
			var d = distance.HsDist(lat, lng, clat, clng)
			truckDistances = append(truckDistances, Distance{truck.Truck, c.Product, d, false})
		}
		routes[truck.Truck] = truckDistances
	}
	return routes
}

func findOptimalRoute(matrix map[string][]Distance) []Distance {
	numRoutes := getLengthOfAnyKey(matrix)
	if numRoutes > 0 {
		optimalRoute := make([]Distance, numRoutes)

		for i := 1; i < numRoutes; i++ {
			var min = math.MaxFloat64
			var dist Distance

			for _, t := range matrix {
				for _, d := range t {
					if !d.isUsed && d.distance < min {
						min = d.distance
						dist = d
					}
				}
			}

			matrix = setUsed(matrix, dist.truck, dist.product)
			optimalRoute = append(optimalRoute, dist)
		}

		return optimalRoute
	}
	return nil

}

func setUsed(routes map[string][]Distance, truck string, product string) map[string][]Distance {
	delete(routes, truck)
	for key, v := range routes {
		for idx, d := range v {
			if d.product == product {
				routes[key][idx].isUsed = true
				continue
			}
		}
	}
	return routes
}

func getLengthOfAnyKey(matrix map[string][]Distance) int {
	for k := range matrix {
		return len(matrix[k])
	}
	return -1
}
