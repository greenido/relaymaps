package main

import (
	"bytes"
	"html/template"
	"log"
)

func KmlToHtml(k *Kml) string {
	pages := GroupByPage(k)

	t, err := template.New("webpage").Parse(tpl)
	if err != nil {
		log.Print("Error rendering the template")
		log.Fatal(err)
	}

	data := struct {
		Pages []*Page
	}{
		Pages: pages,
	}

	var doc bytes.Buffer

	err = t.Execute(&doc, data)
	if err != nil {
		log.Print("Error executing the template")
		log.Fatal(err)
	}

	return doc.String()
}

const tpl = `
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title></title>
<style>
body {
    font-family: Roboto, sans-serif;
    margin: 0;
    padding: 0;
    background: #fdfdfd;
}
h1 {
    color: #ab0000;
    margin-left: 0.6em;
    margin-top: 1em;
}
h2 {
    margin-left: 1em;
    padding-left: 1em;
    position: relative;
    color: #333333;
}
h2:before {
    position: absolute;
    content: '';
    left: 1px;
    top: 30%;
    width: 10px;
    height: 35%;
    background: #ab0000;
}
p {
    margin-left: 3em;
    font-weight: 300;
    position: relative;
    color: #383838;
}
p:before {
    content: '';
    position: absolute;
    left: -12px;
    width: 6px;
    height: 6px;
    background: #6b6b6b;
    border-radius: 3px;
    top: 30%;
}
hr {
    border-top: 1px gray;
    margin: 1em 0 0.5em 0;
    box-shadow: 0 1px 1px #797979;
}
img {
    margin-left: 2em;
    box-shadow: 0 0 8px #bfbfbf;
    margin-top: 0.5em;
}
.pre {
    white-space: pre;
}
</style>
</head>
<body>

{{range .Pages}}
  <hr>
  <h1>Leg {{.Number}}</h1>

  {{with .Start}}
    <h2>From: {{.Pm.Name}}</h2>
    <p>{{.Description}}</p>
    <p>Runner: {{.Runner}}</p>
    {{with .Directions}}<p class="pre">Directions:
{{.}}</p>{{end}}
    {{with .Vans}}<p class="pre">Vans:
{{.}}</p>{{end}}
  {{end}}

  {{with .End}}
    <h2>To: {{.Pm.Name}}</h2>
    <p>{{.Description}}</p>
    {{with .Directions}}<p class="pre">Directions:
{{.}}</p>{{end}}
    {{with .Vans}}<p class="pre">Vans:
{{.}}</p>{{end}}
  {{end}}

  {{with .Leg}}
    <h2>Leg</h2>
    <p>{{.Description}}</p>
    <p>Runner: {{.Runner}}</p>
    {{with .Directions}}<p class="pre">Directions:
{{.}}</p>{{end}}
    {{with .Vans}}<p class="pre">Vans:
{{.}}</p>{{end}}
  {{end}}


  <img src="{{.MapURL}}">
  <img src="{{.ElevationChart}}">
{{end}}

</body>
</html>
`
