package main

import (
	"bytes"
	"os"
	"strings"
	"testing"

	kenyageo "github.com/rmiyoyo/kenya-geo"
)

func TestPagesUpToDate(t *testing.T) {
	rendered, err := render(kenyageo.Default())
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range outputs() {
		old, err := os.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(old, rendered[file]) {
			t.Errorf("%s is out of date; run go run .", file)
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
	if got, want := summit.Dates(), "Thursday 8 to Friday 9 July 2027"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCalendar(t *testing.T) {
	v, err := buildVenue(kenyageo.Default(), summit.Venue, kicc)
	if err != nil {
		t.Fatal(err)
	}
	raw := string(summit.Calendar(v))
	for _, line := range strings.Split(strings.TrimSuffix(raw, "\r\n"), "\r\n") {
		if len(line) > 75 {
			t.Errorf("line is %d bytes, over the 75 the format allows: %q", len(line), line)
		}
	}
	ics := strings.ReplaceAll(raw, "\r\n ", "")
	for _, want := range []string{
		"DTSTART;VALUE=DATE:20270708\r\n",
		"DTEND;VALUE=DATE:20270710\r\n",
		"SUMMARY:InnovationX Commons Africa Summit 2027\r\n",
		"LOCATION:Kenyatta International Conference Centre\\, Harambee Avenue\\, Nairobi\r\n",
	} {
		if !strings.Contains(ics, want) {
			t.Errorf("calendar is missing %q:\n%s", want, ics)
		}
	}
}
