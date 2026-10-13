package main

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

const siteURL = "https://rmiyoyo.github.io/Conference-Website/"

type Session struct {
	Name, Icon, Text string
}

type Speaker struct {
	Name, Role, Bio, Photo string
}

type Slot struct {
	Time, Title, Room string
}

type Day struct {
	Label string
	ISO   string
	Slots []Slot
}

type Partner struct {
	Name, Logo string
}

type Event struct {
	Name     string
	Short    string
	Start    time.Time
	End      time.Time
	Venue    string
	Email    string
	Program  []Session
	Speakers []Speaker
	Partners []Partner
	Schedule [][]Slot
}

func (e Event) Dates() string {
	return e.Start.Format("Monday 2") + " to " + e.End.Format("Monday 2 January 2006")
}

func (e Event) Days() []Day {
	var days []Day
	for i, slots := range e.Schedule {
		d := e.Start.AddDate(0, 0, i)
		days = append(days, Day{d.Format("Monday 2 January"), d.Format("2006-01-02"), slots})
	}
	return days
}

func (e Event) ISOStart() string { return e.Start.Format("2006-01-02") }
func (e Event) ISOEnd() string   { return e.End.Format("2006-01-02") }
func (e Event) Year() int        { return e.Start.Year() }

func (e Event) LogoDates() string {
	return e.Start.Format("2") + "–" + e.End.Format("2 January 2006") + " · Nairobi, Kenya"
}

var summit = Event{
	Name:  "InnovationX Commons Africa Summit",
	Short: "InnovationX Summit",
	Start: time.Date(2027, time.July, 8, 0, 0, 0, 0, time.UTC),
	End:   time.Date(2027, time.July, 9, 0, 0, 0, 0, time.UTC),
	Venue: "Kenyatta International Conference Centre",
	Email: "miyoyo@gmail.com",
	Program: []Session{
		{"Lecture", "images/program/1 lecture.png", "Speakers from across the continent on innovation and technology in Africa."},
		{"Exhibition", "images/program/1 exhibition.png", "Work built on open patents by innovators, hobbyists and hackers."},
		{"Forum", "images/program/1 forum.png", "Put your questions and opinions to experts and leaders in the tech industry."},
		{"Workshop", "images/program/1 workshop.png", "Make something yourself with open source tools instead of only watching."},
		{"Ignite", "images/program/1 ignite.png", "Meet attendees and supporters from across the continent, and stay in touch after the summit."},
	},
	Speakers: []Speaker{
		{"James Smith", "Chief Technology Officer, Innovate Tech", "James leads a team of engineers and researchers building new products. He has spent over ten years in software development.", "images/keynote/naledi.png"},
		{"Sarah Johnson", "Cybersecurity Expert, SecureNet", "Sarah protects organisations from cyber threats and speaks at industry conferences about good security practice.", "images/keynote/lwazi.png"},
		{"Michael Lee", "Data Scientist, DataWise", "Michael works on machine learning and analytics, and on getting businesses to trust what their data says.", "images/keynote/fadumo.png"},
		{"Emily Chen", "Product Manager, Tech Innovations", "Emily has a background in product management and user experience design, and cares about products people actually use.", "images/keynote/yemisi.png"},
		{"Robert Brown", "Blockchain Developer, Crypto Solutions", "Robert builds systems for secure transactions and has worked on several large blockchain projects.", "images/keynote/hicham.png"},
		{"Linda Garcia", "AI Researcher, Future AI Labs", "Linda develops AI methods for hard problems and takes part in work on responsible AI.", "images/keynote/doreen.png"},
	},
	Schedule: [][]Slot{
		{
			{"08:00", "Registration and coffee", "Main foyer"},
			{"09:00", "Opening keynote: James Smith on building research teams that ship", "Plenary Hall"},
			{"10:30", "Lecture: Sarah Johnson on security for small teams", "Plenary Hall"},
			{"11:30", "Exhibition opens", "Tsavo Ballroom"},
			{"13:00", "Lunch", "Main foyer"},
			{"14:00", "Workshops", "Amboseli Hall"},
			{"16:00", "Forum: Michael Lee and Emily Chen on turning data into products", "Plenary Hall"},
			{"18:00", "Ignite evening", "Rooftop terrace"},
		},
		{
			{"08:30", "Coffee", "Main foyer"},
			{"09:00", "Keynote: Linda Garcia on responsible AI", "Plenary Hall"},
			{"10:30", "Lecture: Robert Brown on blockchain beyond money", "Plenary Hall"},
			{"11:30", "Exhibition", "Tsavo Ballroom"},
			{"13:00", "Lunch", "Main foyer"},
			{"14:00", "Workshops", "Amboseli Hall"},
			{"16:00", "Open forum with all six speakers", "Plenary Hall"},
			{"17:30", "Closing", "Plenary Hall"},
		},
	},
	Partners: []Partner{
		{"YouTube", "images/partners/youtube.png"},
		{"Meta", "images/partners/facebook.png"},
		{"Microsoft", "images/partners/windows.png"},
		{"Ford", "images/partners/ford.png"},
		{"Adidas", "images/partners/grand.png"},
	},
}

func (e Event) Calendar(v Venue) []byte {
	lines := []string{
		"BEGIN:VCALENDAR",
		"VERSION:2.0",
		"PRODID:-//" + e.Short + "//EN",
		"CALSCALE:GREGORIAN",
		"BEGIN:VEVENT",
		"UID:" + e.ISOStart() + "@innovationx-summit",
		"DTSTAMP:" + e.Start.Format("20060102") + "T000000Z",
		"DTSTART;VALUE=DATE:" + e.Start.Format("20060102"),
		"DTEND;VALUE=DATE:" + e.End.AddDate(0, 0, 1).Format("20060102"),
		"SUMMARY:" + icsText(fmt.Sprintf("%s %d", e.Name, e.Year())),
		"LOCATION:" + icsText(v.Name+", Harambee Avenue, Nairobi"),
		fmt.Sprintf("GEO:%.4f;%.4f", v.Lat, v.Lng),
		"URL:" + siteURL,
		"END:VEVENT",
		"END:VCALENDAR",
	}
	var b strings.Builder
	for _, l := range lines {
		for len(l) > 75 {
			cut := 75
			for cut > 0 && !utf8.RuneStart(l[cut]) {
				cut--
			}
			b.WriteString(l[:cut] + "\r\n ")
			l = l[cut:]
		}
		b.WriteString(l + "\r\n")
	}
	return []byte(b.String())
}

func icsText(s string) string {
	return strings.NewReplacer(`\`, `\\`, ";", `\;`, ",", `\,`).Replace(s)
}

type place struct {
	Type    string `json:"@type"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type eventData struct {
	Context   string `json:"@context"`
	Type      string `json:"@type"`
	Name      string `json:"name"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Mode      string `json:"eventAttendanceMode"`
	Status    string `json:"eventStatus"`
	URL       string `json:"url"`
	Location  place  `json:"location"`
}

func (e Event) StructuredData(v Venue) eventData {
	return eventData{
		Context:   "https://schema.org",
		Type:      "Event",
		Name:      fmt.Sprintf("%s %d", e.Name, e.Year()),
		StartDate: e.ISOStart(),
		EndDate:   e.ISOEnd(),
		Mode:      "https://schema.org/OfflineEventAttendanceMode",
		Status:    "https://schema.org/EventScheduled",
		URL:       siteURL,
		Location:  place{"Place", v.Name, "Harambee Avenue, Nairobi, Kenya"},
	}
}

func (e Event) URL() string { return siteURL }
