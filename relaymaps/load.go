package main

// Will be useful to encode polylines.
// https://github.com/twpayne/go-polyline

import (
	"archive/zip"
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func ReadKML(reader io.Reader) (*Kml, error) {
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Printf("Error reading kml file")
		return nil, err
	}

	// A bit hacky, but if there is a "<NetworkLink>" in
	// this file, we will instead try again with that file.
	if bytes.Contains(content, []byte("<NetworkLink>")) {
		from := bytes.Index(content, []byte("<href>"))
		to := bytes.Index(content, []byte("</href>"))
		url := content[from+6 : to]
		log.Printf("Following <NetworkLink> to: %s", url)
		return FetchAndReadKMZ(string(url))
	}

	doc, err := Unmarshal(content)
	if err != nil {
		log.Printf("Error unmarshaling kml file")
		return nil, err
	}
	return doc, nil
}

func FetchAndReadKMZ(url string) (*Kml, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching %s", url)
		return nil, err
	}
	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response for %s", url)
		return nil, err
	}

	reader := bytes.NewReader(all)
	r, err := zip.NewReader(reader, int64(len(all)))
	if err != nil {
		log.Printf("Error reading zipped response for %s", url)
		return nil, err
	}
	return ReadKMZ(r)
}

func OpenKMZ(filename string) (*Kml, error) {
	r, err := zip.OpenReader(filename)
	if err != nil {
		log.Printf("Error opening %s", filename)
		return nil, err
	}
	defer r.Close()
	return ReadKMZ(&r.Reader)
}

func ReadKMZ(r *zip.Reader) (*Kml, error) {
	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		if strings.HasSuffix(f.Name, ".kml") {
			rc, err := f.Open()
			defer rc.Close()
			if err != nil {
				log.Printf("Error reading the kmz file")
				return nil, err
			}
			return ReadKML(rc)
		}
	}
	return nil, errors.New("No kml found in kmz")
}
