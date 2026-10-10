package main

import "time"

type Session struct {
	Name, Icon, Text string
}

type Speaker struct {
	Name, Role, Bio, Photo string
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
}

func (e Event) Dates() string {
	return e.Start.Format("Monday 2") + " to " + e.End.Format("Monday 2 January 2006")
}

func (e Event) ISOStart() string { return e.Start.Format("2006-01-02") }
func (e Event) ISOEnd() string   { return e.End.Format("2006-01-02") }
func (e Event) Year() int        { return e.Start.Year() }

var summit = Event{
	Name:  "InnovationX Commons Africa Summit",
	Short: "InnovationX Summit",
	Start: time.Date(2025, time.July, 10, 0, 0, 0, 0, time.UTC),
	End:   time.Date(2025, time.July, 11, 0, 0, 0, 0, time.UTC),
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
		{"Michael Lee", "Data Scientist, DataWise", "Michael works on machine learning and analytics, and on getting businesses to trust what their data says.", "images/keynote/yemisi.png"},
		{"Emily Chen", "Product Manager, Tech Innovations", "Emily has a background in product management and user experience design, and cares about products people actually use.", "images/keynote/hicham.png"},
		{"Robert Brown", "Blockchain Developer, Crypto Solutions", "Robert builds systems for secure transactions and has worked on several large blockchain projects.", "images/keynote/doreen.png"},
		{"Linda Garcia", "AI Researcher, Future AI Labs", "Linda develops AI methods for hard problems and takes part in work on responsible AI.", "images/keynote/fadumo.png"},
	},
	Partners: []Partner{
		{"YouTube", "images/partners/youtube.png"},
		{"Meta", "images/partners/facebook.png"},
		{"Microsoft", "images/partners/windows.png"},
		{"Ford", "images/partners/ford.png"},
		{"Adidas", "images/partners/grand.png"},
	},
}
