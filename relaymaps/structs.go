package main

// Those structs were mostly taken from
// https://github.com/joshuamorris3/gokml-gx
// I copied them here so I could quickly add what was missing (mostly just
// LineString actually).

import (
	"encoding/xml"
)

func Unmarshal(doc []byte) (*Kml, error) {
	kml := &Kml{}
	error := xml.Unmarshal(doc, kml)
	return kml, error
}

type Kml struct {
	XMLName    xml.Name     `xml:"kml"`
	Namespace  string       `xml:"xmlns,attr"`
	Document   *Document    `xml:"Document"`
	Placemarks []*Placemark `xml:"Placemark"`
	Folders    []*Folder    `xml:"Folder"`
}

type Document struct {
	XMLName     xml.Name     `xml:"Document"`
	Id          string       `xml:"id,attr,omitempty"`
	Name        string       `xml:"name,omitempty"`
	Visibility  int          `xml:"visibility,omitempty"`
	Open        int          `xml:"open,omitempty"`
	Address     string       `xml:"address,omitempty"`
	PhoneNumber string       `xml:"phoneNumber,omitempty"`
	Description string       `xml:"description,omitempty"`
	DocStyle    []*Style     `xml:"Style"`
	Placemarks  []*Placemark `xml:"Placemark"`
	Folders     []*Folder    `xml:"Folder"`
}

type Placemark struct {
	XMLName     xml.Name      `xml:"Placemark"`
	Id          string        `xml:"id,attr,omitempty"`
	Name        string        `xml:"name,omitempty"`
	Description string        `xml:"description,omitempty"`
	StyleUrl    string        `xml:"styleUrl,omitempty"`
	Coordinates string        `xml:"Point>coordinates"`
	LineString  string        `xml:"LineString>coordinates"`
	Extended    *ExtendedData `xml:"ExtendedData"`
}

type Folder struct {
	XMLName     xml.Name     `xml:"Folder"`
	Id          string       `xml:"id,attr,omitempty"`
	Name        string       `xml:"name,omitempty"`
	Visibility  int          `xml:"visibility,omitempty"`
	Open        int          `xml:"open,omitempty"`
	Address     string       `xml:"address,omitempty"`
	PhoneNumber string       `xml:"phoneNumber,omitempty"`
	Description string       `xml:"description,omitempty"`
	Placemarks  []*Placemark `xml:"Placemark"`
	Folders     []*Folder    `xml:"Folder"`
}

type Style struct {
	XMLName xml.Name   `xml:"Style"`
	Id      string     `xml:"id,attr,omitempty"`
	Line    *LineStyle `xml:"LineStyle"`
	Icon    *IconStyle `xml:"IconStyle"`
}

type IconStyle struct {
	XMLName xml.Name `xml:"IconStyle"`
	Scale   string   `xml:"scale,omitempty"`
	Heading string   `xml:"heading,omitempty"`
	Href    string   `xml:"Icon>href,omitempty"`
}

type LineStyle struct {
	XMLName xml.Name `xml:"LineStyle"`
	Color   string   `xml:"color,omitempty"`
	Width   string   `xml:"width,omitempty"`
}

type ExtendedData struct {
	XMLName xml.Name `xml:"ExtendedData"`
	Datas   []*Data  `xml:"Data"`
}

type Data struct {
	XMLName xml.Name `xml:"Data"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value"`
}
