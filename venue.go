package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	kenyageo "github.com/rmiyoyo/kenya-geo"
)

var kicc = kenyageo.LatLng{Lat: -1.2884, Lng: 36.8233}

type Venue struct {
	Name         string
	Ward         string
	Constituency string
	County       string
	CountyCode   int
	Lat, Lng     float64
	Counties     []TravelCounty
}

type TravelCounty struct {
	Code         int
	Name         string
	Headquarters string
	Km           int
	Borders      bool
	Host         bool
}

func (v Venue) MapLink() string {
	return fmt.Sprintf("https://www.openstreetmap.org/?mlat=%.4f&mlon=%.4f#map=17/%.4f/%.4f", v.Lat, v.Lng, v.Lat, v.Lng)
}

func buildVenue(geo *kenyageo.Data, name string, at kenyageo.LatLng) (Venue, error) {
	w, ok := geo.WardAt(at.Lat, at.Lng)
	if !ok {
		return Venue{}, fmt.Errorf("no ward at %v", at)
	}
	v := Venue{
		Name:         name,
		Ward:         titleCase(w.Name),
		Constituency: w.Constituency,
		County:       w.CountyName,
		CountyCode:   w.CountyCode,
		Lat:          at.Lat,
		Lng:          at.Lng,
	}
	borders := map[int]bool{}
	for _, c := range geo.NeighbouringCounties(w.CountyCode) {
		borders[c.Code] = true
	}
	for _, c := range geo.Counties() {
		v.Counties = append(v.Counties, TravelCounty{
			Code:         c.Code,
			Name:         c.Name,
			Headquarters: c.Headquarters,
			Km:           int(math.Round(haversineKm(c.Centroid, at))),
			Borders:      borders[c.Code],
			Host:         c.Code == w.CountyCode,
		})
	}
	sort.Slice(v.Counties, func(i, j int) bool { return v.Counties[i].Name < v.Counties[j].Name })
	return v, nil
}

func haversineKm(a, b kenyageo.LatLng) float64 {
	const r = 6371.0
	rad := math.Pi / 180
	dLat := (b.Lat - a.Lat) * rad
	dLng := (b.Lng - a.Lng) * rad
	h := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(a.Lat*rad)*math.Cos(b.Lat*rad)*math.Sin(dLng/2)*math.Sin(dLng/2)
	return 2 * r * math.Asin(math.Sqrt(h))
}

func titleCase(s string) string {
	words := strings.Fields(strings.ToLower(s))
	for i, w := range words {
		words[i] = strings.ToUpper(w[:1]) + w[1:]
	}
	return strings.Join(words, " ")
}
