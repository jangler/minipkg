package webui

import "html/template"

// PageTemplate is an HTML template for a complete page.
var PageTemplate = template.Must(template.New("page").Parse(`
<!DOCTYPE html>
<html>
<head>
<title>{{.Title}}</title>
<style>
a,h1,h2,h3,h4,h5,h6 { color: #375eab; }
body { background-color: white; font-family: sans-serif; padding: 0.5rem; }
.title { background-color: #e0ebf5; padding: 0.5rem; }
.title > h2 { margin: 0; }
.content { padding: 1rem; }
html { background-color: whitesmoke; }
p { color: #222222; }
</style>
</head>
<body>
<div class="title"><h2>{{.Title}}</h2></div>
<div class="content">
{{.Content}}
</div>
</body>
</html>
`))

// PageData is a valid data argument for PageTemplate.
type PageData struct {
	Title   string
	Content template.HTML
}
