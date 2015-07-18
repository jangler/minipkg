package webui

import (
	"fmt"
	"net"
	"net/http"

	"github.com/jangler/minipkg/writers"
	"github.com/pkg/browser"
)

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
