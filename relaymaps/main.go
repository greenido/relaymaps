package main

import (
	"flag"
	"io/ioutil"
	"log"
)

var (
	kmz = flag.String("kmz", "", "KMZ file")
	out = flag.String("out", "", "HTML output file")
)

func main() {
	flag.Parse()
	file, err := OpenKMZ(*kmz)
	if err != nil {
		log.Fatal(err)
	} else if file == nil {
		log.Fatal("Returned a nil kml file")
	}

	html := KmlToHtml(file)
	ioutil.WriteFile(*out, []byte(html), 0x0666)
}
