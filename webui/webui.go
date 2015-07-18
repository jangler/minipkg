package webui

import (
	"fmt"
	"html/template"
	"net"
	"net/http"

	"github.com/jangler/minipkg/writers"
	"github.com/pkg/browser"
)

// Template is a template.Template that takes a Context as its data argument.
var Template = template.Must(template.New("tmpl").Parse(`
<!DOCTYPE html>
<html>
<head>
<title>{{.Title}}</title>
<style>
body { font-family: sans-serif; padding: 0.5rem; }
div.title { background-color: #e0ebf5; padding: 0.5rem; }
div.content { padding: 0.5rem; }
h1,h2,h3,h4,h5,h6 { color: #375eab; margin: 0.5rem; }
p { color: #222222; margin: 0.5rem 1.5rem 0.5rem 1.5rem; }
</style>
</head>
<body>
<div class="title"><h1>{{.Title}}</h1></div>
<div class="content">{{.Content}}</div>
</body>
</html>
`))

// Context is a valid data argument for Template.Execute().
type Context struct {
	Title   string
	Content template.HTML
}

func getOpenPort() string {
	port := 8080
	for {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err == nil {
			listener.Close()
			break
		}
		port++
	}
	return fmt.Sprintf(":%d", port)
}

// Start starts the HTTP server and opens the application page in the user's
// web browser, or returns an error.
func Start() error {
	done := make(chan error)
	addr := getOpenPort()
	go func() {
		done <- http.ListenAndServe(addr, nil)
	}()
	var d writers.Discarder
	browser.Stdout, browser.Stderr = d, d
	browser.OpenURL(fmt.Sprintf("http://localhost%s/", addr))
	return <-done
}
