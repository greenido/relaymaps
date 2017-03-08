package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/vincent-petithory/dataurl"
	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

func meterToFeet(m float64) float64 {
	return m * 3.28084
}

func MakeElevationPlot(p *Page) template.URL {
	xys := GetElevations(p)

	xs := make([]float64, len(xys))
	ys := make([]float64, len(xys))
	num := len(xys)
	for i, pt := range xys {
		xs[i] = float64(i) / float64(num-1)
		ys[i] = meterToFeet(pt.Elevation)
	}

	log.Printf("Xs = %#v", xs)
	log.Printf("Ys = %#v", ys)

	graph := chart.Chart{
		Width:  500,
		Height: 200,
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		XAxis: chart.XAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: xs,
				YValues: ys,
				Style: chart.Style{
					Show:        true,                           //note; if we set ANY other properties, we must set this to true.
					StrokeColor: drawing.ColorBlue,               // will supercede defaults
					FillColor:   drawing.ColorBlue.WithAlpha(64), // will supercede defaults
				},
			},
			chart.AnnotationSeries{
				Annotations: []chart.Value2{
					{
						XValue: 1.0,
						YValue: ys[num-1],
						Label:  fmt.Sprintf("%.0f", ys[num-1]),
					},
				},
			},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		log.Fatalf("Error rendering elevations chart: %s", err)
	}

	u := dataurl.New(buffer.Bytes(), "image/png")
	return template.URL(u.String())
}

var limiter = time.Tick(time.Millisecond * 200)

func GetElevations(p *Page) []ElevationPoint {
	params := url.Values{}
	params.Add("path", "enc:"+MakePolyline(p.Leg.Pm.LineString))
	params.Add("samples", "60")

	// May need an API Key here
	//params.Add("key", "")

	url := ("https://maps.googleapis.com/maps/api/elevation/json?" +
		params.Encode())
	log.Printf("URL: %s", url)

	<-limiter
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error getting elevations for path: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading elevations response: %s", err)
	}

	elevationResp := &ElevationResp{}
	err = json.Unmarshal(body, &elevationResp)
	if err != nil {
		log.Fatalf("Error parsing elevations response: %s", err)
	}
	if elevationResp.Status != "OK" {
		log.Printf("%#v", elevationResp)
		log.Fatalf("Got a non-OK status from elevation req: %s",
			elevationResp.Status)
	}

	return elevationResp.Results
}

type Xy struct {
	X, Y float64
}

type ElevationResp struct {
	Results []ElevationPoint
	Status  string
}

type ElevationPoint struct {
	Elevation  float64
	Location   Location
	Resulution float64
}

type Location struct {
	Lat float64
	Lng float64
}
