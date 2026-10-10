package main

import (
	"bytes"
	"os"
	"testing"

	kenyageo "github.com/rmiyoyo/kenya-geo"
)

func TestPagesUpToDate(t *testing.T) {
	rendered, err := render(kenyageo.Default())
	if err != nil {
		t.Fatal(err)
	}
	for _, p := range pages {
		old, err := os.ReadFile(p.File)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(old, rendered[p.File]) {
			t.Errorf("%s is out of date; run go run .", p.File)
		}
	}
}

func TestVenue(t *testing.T) {
	v, err := buildVenue(kenyageo.Default(), summit.Venue, kicc)
	if err != nil {
		t.Fatal(err)
	}
	if v.Ward != "Nairobi Central" || v.Constituency != "Starehe" || v.County != "Nairobi County" {
		t.Errorf("venue is in %s, %s, %s", v.Ward, v.Constituency, v.County)
	}
	if len(v.Counties) != 47 {
		t.Fatalf("got %d counties, want 47", len(v.Counties))
	}
	hosts, borders := 0, 0
	for _, c := range v.Counties {
		if c.Host {
			hosts++
		}
		if c.Borders {
			borders++
		}
		if c.Km < 0 || c.Km > 1000 {
			t.Errorf("%s is %d km away", c.Name, c.Km)
		}
	}
	if hosts != 1 || borders != 3 {
		t.Errorf("got %d host and %d bordering counties, want 1 and 3", hosts, borders)
	}
}

func TestHaversine(t *testing.T) {
	mombasa := kenyageo.LatLng{Lat: -4.0435, Lng: 39.6682}
	if km := haversineKm(kicc, mombasa); km < 430 || km > 450 {
		t.Errorf("KICC to Mombasa is %.0f km, want about 440", km)
	}
}

func TestDates(t *testing.T) {
	if got, want := summit.Dates(), "Thursday 10 to Friday 11 July 2025"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
