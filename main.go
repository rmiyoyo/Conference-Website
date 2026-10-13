package main

import (
	"bytes"
	"embed"
	"flag"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"

	kenyageo "github.com/rmiyoyo/kenya-geo"
)

//go:embed templates/*.tmpl
var templates embed.FS

type page struct {
	File, Template, Title, Current string
}

var pages = []page{
	{"index.html", "index.tmpl", "", "home"},
	{"about_page.html", "about.tmpl", "About", "about"},
}

type view struct {
	Page  page
	Event Event
	Venue Venue
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("site", flag.ContinueOnError)
	fs.SetOutput(stderr)
	out := fs.String("out", ".", "directory to write the pages to")
	check := fs.Bool("check", false, "fail if the pages in -out are not up to date")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	rendered, err := render(kenyageo.Default())
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	stale := 0
	for _, file := range outputs() {
		path := filepath.Join(*out, file)
		if *check {
			old, err := os.ReadFile(path)
			if err != nil || !bytes.Equal(old, rendered[file]) {
				fmt.Fprintf(stderr, "%s is out of date; run go run .\n", path)
				stale++
			}
			continue
		}
		if err := os.WriteFile(path, rendered[file], 0o644); err != nil {
			fmt.Fprintln(stderr, err)
			return 1
		}
		fmt.Fprintln(stdout, "wrote", path)
	}
	if stale > 0 {
		return 1
	}
	return 0
}

func render(geo *kenyageo.Data) (map[string][]byte, error) {
	venue, err := buildVenue(geo, summit.Venue, kicc)
	if err != nil {
		return nil, err
	}
	out := map[string][]byte{}
	for _, p := range pages {
		t, err := template.ParseFS(templates, "templates/layout.tmpl", "templates/"+p.Template)
		if err != nil {
			return nil, err
		}
		var buf bytes.Buffer
		if err := t.ExecuteTemplate(&buf, "layout", view{p, summit, venue}); err != nil {
			return nil, fmt.Errorf("%s: %w", p.File, err)
		}
		out[p.File] = buf.Bytes()
	}
	out[calendarFile] = summit.Calendar(venue)
	return out, nil
}

const calendarFile = "summit.ics"

func outputs() []string {
	files := []string{calendarFile}
	for _, p := range pages {
		files = append(files, p.File)
	}
	return files
}
