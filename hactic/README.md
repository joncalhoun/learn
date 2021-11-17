# A Hacktastic Static Site Generator

In a `main.go`:

```go
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
```

Then define some `gohtml` files.

```html
<!-- index.gohtml -->
{{define "index"}}
<html>
<head>
  <!--
    It is pretty easy to add this to the build pipeline, but not shown here.
  -->
  <link href="https://unpkg.com/tailwindcss@^2/dist/tailwind.min.css" rel="stylesheet">
</head>
<body>
	{{quote "Mat Ryer" "/img/mat-avatar.jpg" "Jon is brilliant. A modern Albert Einstein!"}}
	{{quote "Jon Calhoun" "/img/jon-avatar.jpg" "Don't believe a word Mat says." "He likes to yank my chain."}}
</body>
</html>
{{end}}
```

```html
<!-- quote.gohtml -->
{{define "quote"}}
<blockquote class="px-12 bg-gray-100 rounded-lg m-12 flex space-x-8">
	<div class="text-center py-8">
    {{if .Img}}<img class="w-32 h-32 rounded-full" src="{{.Img}}"/>{{else}}<svg>faking this...</svg>{{end}}
    <div>{{.Who}}</div>
  </div>
	<div class="flex-grow text-lg py-12">
		{{range .Quote}}<p>{{.}}</p>{{end}}
	</div>
</blockquote>
{{end}}
```

Setup a `modd.conf` file (see [modd](https://github.com/cortesi/modd))

```confg
**/*.go !**/*_test.go **/*.gohtml {
  prep: go run *.go
}

touch-to-reload-simple-server.txt {
  daemon +sigterm: python -m SimpleHTTPServer 1313
}
```

And run it.

```bash
modd
```

Visit <localhost:1313> for your static site.


