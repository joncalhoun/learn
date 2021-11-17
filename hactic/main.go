package main

import (
	"html/template"
	"os"
	"strings"
)

func main() {
	var tpl *template.Template
	tpl = template.Must(template.New("").Funcs(template.FuncMap{
		// Add custom functions for reusable components with variable data. Super hacky. 😂
		"quote": func(who, img string, quote ...string) template.HTML {
			data := struct {
				Who   string
				Img   string
				Quote []string
			}{
				Who:   who,
				Img:   img,
				Quote: quote,
			}
			var sb strings.Builder
			tpl.ExecuteTemplate(&sb, "quote", data)
			return template.HTML(sb.String())
		},
	}).ParseGlob("*.gohtml"))

	// This is how we write to an index.html file.
	f, err := os.Create("index.html")
	if err != nil {
		panic(err)
	}
	err = tpl.ExecuteTemplate(f, "index", nil)
	if err != nil {
		panic(err)
	}
}
