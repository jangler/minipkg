package webui

import "html/template"

// TableTemplate is an HTML template for a table.
var TableTemplate = template.Must(template.New("table").Parse(`
<table>
{{if .Caption}}<caption>{{.Caption}}</caption>{{end}}
{{if .Header}}
<thead><tr>
{{range .Header}}
<th>{{.}}</th>
{{end}}
</tr></thead>
{{end}}
{{if .Footer}}
<tfoot><tr>
{{range .Footer}}
<td>{{.}}</td>
{{end}}
</tr></tfoot>
{{end}}
{{if .Body}}
<tbody>
{{range .Body}}
<tr>
{{range .}}
<td>{{.}}</td>
{{end}}
</tr>
{{end}}
</tbody>
{{end}}
</table>
`))

// TableData is a valid data argument for TableTemplate.
type TableData struct {
	Caption string
	Header  []string
	Footer  []string
	Body    [][]string
}
