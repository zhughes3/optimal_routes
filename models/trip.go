package models

import (
	"fmt"
	"github.com/zhughes3/optimal_routes/distance"
	"math"
)

type TripList struct {
	routes []trip
}

type trip struct {
	truck string
	product string
	distance float64
}

type tripInfo struct {
	truck          string
	product        string
	pickUpDistance float64
	deliveryDistance float64
	isUsed         bool
}

func (tl TripList) addTrip(t tripInfo) TripList {
	tl.routes = append(tl.routes, trip{t.truck,t.product, t.pickUpDistance + t.deliveryDistance})
	return tl
}

func (tl TripList) Print() {
	for _, trip := range tl.routes {
		fmt.Println(trip.truck, "will deliver", trip.product, "for a total distance of", trip.distance)
	}
}

func InitMatrix() map[string][]tripInfo {
	trucks := ParseTrucks("trucks.csv")
	cargo := ParseCargos("cargo.csv")

	routes := make(map[string][]tripInfo)
	for _, truck := range trucks {
		var truckDistances []tripInfo
		var lat, lng = truck.Lat, truck.Lng
		for _, c := range cargo {
			var clat, clng = c.SrcLat, c.SrcLng
			var dLat, dLng = c.DstLat, c.DstLng
			var pickupDist = distance.HsDist(lat, lng, clat, clng)
			var deliveryDist = distance.HsDist(clat, clng, dLat, dLng)
			truckDistances = append(truckDistances, tripInfo{truck.Truck, c.Product, pickupDist, deliveryDist, false})
		}
		routes[truck.Truck] = truckDistances
	}
	return routes
}

func FindOptimalRoutes(matrix map[string][]tripInfo) TripList {
	numRoutes := getLengthOfAnyKey(matrix)
	if numRoutes > 0 {

		var trips TripList

		for i := 0; i < numRoutes; i++ {
			var min = math.MaxFloat64
			var dist tripInfo

			for _, t := range matrix {
				for _, d := range t {
					if !d.isUsed && d.pickUpDistance < min {
						min = d.pickUpDistance
						dist = d
					}
				}
			}

			matrix = setUsed(matrix, dist.truck, dist.product)

			trips = trips.addTrip(dist)
		}

		return trips
	}
	return TripList{}

}

func setUsed(routes map[string][]tripInfo, truck string, product string) map[string][]tripInfo {
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

func getLengthOfAnyKey(matrix map[string][]tripInfo) int {
	for k := range matrix {
		return len(matrix[k])
	}
	return -1
}
