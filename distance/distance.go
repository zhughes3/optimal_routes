package distance

import "math"

// All of the code below has been adapted from // adapted from https://rosettacode.org/wiki/Haversine_formula#Go
// with minimal refactoring

func haversine(θ float64) float64 {
	return .5 * (1 - math.Cos(θ))
}

type pos struct {
	lat float64 // latitude, radians
	lng float64 // longitude, radians
}

func degPos(lat, lon float64) pos {
	return pos{lat * math.Pi / 180, lon * math.Pi / 180}
}

const rEarth = 6372.8 // km

func HsDist(lat1, lng1, lat2, lng2 float64) float64 {
	var p1 = degPos(lat1, lng1)
	var p2 = degPos(lat2, lng2)
	return 2 * rEarth * math.Asin(math.Sqrt(haversine(p2.lat-p1.lat)+
		math.Cos(p1.lat)*math.Cos(p2.lat)*haversine(p2.lng-p1.lng)))
}