package webui

import (
	"html/template"
	"net"
	"net/http"
	"testing"
	"time"
)

func TestGetOpenPort(t *testing.T) {
	if got, want := getOpenPort(), ":8080"; got != want {
		t.Errorf("getOpenPort(0) == %#v; want %#v", got, want)
	}
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		t.Fatalf("%v", err)
	}
	if got, want := getOpenPort(), ":8081"; got != want {
		t.Errorf("getOpenPort(0) == %#v; want %#v", got, want)
	}
	listener.Close()
}

func TestLaunch(t *testing.T) {
	done := make(chan struct{})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Template.Execute(w, Context{
			Title:   "Hello, world!",
			Content: template.HTML("<p>Hello, world!</p>"),
		})
		done <- struct{}{}
	})
	go Start()
	<-done
	time.Sleep(time.Second / 10)
}
