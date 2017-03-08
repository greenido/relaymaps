package main

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"html/template"

	line "github.com/twpayne/go-polyline"
)

type Page struct {
	Number     int
	Start, End *ParsedPlacemark
	Leg        *ParsedPlacemark

	MapURL string

	// A PNG image as bytes.
	ElevationChart template.URL
}

type ParsedPlacemark struct {
	Pm *Placemark
	// This stuff comes from Extended Data.
	Description string
	Runner      string
	URL         string
	Directions  string
	Vans        string
}

// Somehow, kml stores as lon,lat, we want lat,lon...
func ParseCoord(c string) string {
	parts := strings.Split(strings.TrimSpace(c), ",")
	return strings.Join([]string{parts[1], parts[0]}, ",")
}

func MakePolyline(c string) string {
	pts := strings.Split(c, "\n")

	var coords = [][]float64{}
	for _, p := range pts {
		parts := strings.Split(strings.TrimSpace(p), ",")
		if len(parts) >= 2 {
			lat, _ := strconv.ParseFloat(parts[1], 64)
			lon, _ := strconv.ParseFloat(parts[0], 64)
			coords = append(coords, []float64{lat, lon})
		}
	}
	return string(line.EncodeCoords(coords))
}

func MakeMapURL(p *Page) string {
	params := url.Values{}

	// Markers
	markers := ""
	markers += "&markers=label:S|" + ParseCoord(p.Start.Pm.Coordinates)
	markers += "&markers=label:E|" + ParseCoord(p.End.Pm.Coordinates)

	params.Add("size", "600x400")

	//&path=weight:3%7Ccolor:orange%7Cenc:
	params.Add("path", "weight:5|color:blue|enc:"+
		MakePolyline(p.Leg.Pm.LineString))

	return ("https://maps.googleapis.com/maps/api/staticmap?" +
		params.Encode() +
		markers)
}

func ParsePlacemark(p *Placemark) *ParsedPlacemark {
	res := &ParsedPlacemark{
		Pm: p,
	}

	if p.Extended != nil {
		for _, d := range p.Extended.Datas {
			switch strings.ToLower(d.Name) {
			case "description":
				res.Description = d.Value
			case "runner #":
				res.Runner = d.Value
			case "url":
				res.URL = d.Value
			case "directions":
				res.Directions = d.Value
			case "vans":
				res.Vans = d.Value
			}
		}
	}
	return res
}

func FindPlacemark(pms []*Placemark, prefix string) *ParsedPlacemark {
	for _, p := range pms {
		if matched, _ := regexp.MatchString(
			"^"+prefix+"($|[ ,.])", p.Name); matched {
			return ParsePlacemark(p)
		}
	}
	return nil
}

func ExchangeRE(n int) string {
	return fmt.Sprintf("Exch(.|ange) %d", n)
}

func NewPage(pms []*Placemark, leg int) *Page {
	// Find the start:
	startPrefix := ExchangeRE(leg - 1)
	endPrefix := ExchangeRE(leg)
	legPrefix := fmt.Sprintf("Leg %d", leg)

	if leg == 1 {
		startPrefix = "Start"
	}
	if leg == NumPages {
		endPrefix = "Finish"
	}

	page := &Page{
		Number: leg,
		Start:  FindPlacemark(pms, startPrefix),
		End:    FindPlacemark(pms, endPrefix),
		Leg:    FindPlacemark(pms, legPrefix),
	}

	page.MapURL = MakeMapURL(page)
	page.ElevationChart = MakeElevationPlot(page)
	return page
}

const NumPages = 36

func GroupByPage(k *Kml) []*Page {
	pms := AllPlacemarks(k)
	pages := make([]*Page, NumPages)
	for i := 1; i <= NumPages; i++ {
		pages[i-1] = NewPage(pms, i)
	}
	return pages
}

func AllPlacemarks(k *Kml) []*Placemark {
	placemarks := make([]*Placemark, 0)
	placemarks = append(placemarks, k.Placemarks...)
	placemarks = append(placemarks, k.Document.Placemarks...)
	for _, folder := range k.Folders {
		placemarks = append(placemarks, CollectPlacemarks(folder)...)
	}
	for _, folder := range k.Document.Folders {
		placemarks = append(placemarks, CollectPlacemarks(folder)...)
	}
	return placemarks
}

func CollectPlacemarks(f *Folder) []*Placemark {
	p := make([]*Placemark, 0)
	for _, sub := range f.Folders {
		p = append(p, CollectPlacemarks(sub)...)
	}
	p = append(p, f.Placemarks...)
	return p
}
