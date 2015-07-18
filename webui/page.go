package webui

import "html/template"

// PageTemplate is an HTML template for a complete page.
var PageTemplate = template.Must(template.New("page").Parse(`
<!DOCTYPE html>
<html>
<head>
<title>{{.Title}}</title>
<meta name="viewport" content="width=device-width, initial-scale=1">
<style>
a { text-decoration: none; }
a:hover { text-decoration: underline; }
a,h1,h2,h3,h4,h5,h6 { color: #38468d; }
body { background-color: #ffffff; font-family: sans-serif; margin: 0; }
.title { background-color: #dee1f6; padding: 1rem; }
.title > h2 { margin: 0; }
.content { padding: 1rem; }
html { background-color: #e2e2e2; }
p { color: #2f2f2f; }
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
