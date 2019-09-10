package server

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gobuffalo/packr/v2"
	"github.com/tcd/shu/internal/shu"
)

var tmpl *template.Template

func init() {
	box := packr.New("name", "./static/template")

	index, err := box.FindString("index.gohtml") // Living on the edge, not handling this error.
	tmpl, err = parseMany(map[string]string{
		"index": index,
	})
	if err != nil {
		log.Fatal(err)
	}
}

// ServeToPort spins up a web interface for shu on a given port.
func ServeToPort(portNumber string) {
	setupCloseHandler()

	box := packr.New("someBoxName", "./static/assets")

	http.HandleFunc("/", indexHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(box)))

	fmt.Printf("Web Server is available at http://localhost:%s/\n", portNumber)
	err := http.ListenAndServe(":"+portNumber, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// Serve spins up a web interface for shu.
// An open port is used and printed to stdout.
func Serve() {
	setupCloseHandler()

	box := packr.New("someBoxName", "./static/assets")

	http.HandleFunc("/", indexHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(box)))

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Web Server is available at http://localhost:%v/\n", listener.Addr().(*net.TCPAddr).Port)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	shoes, err := shu.GetIssues()
	if err != nil {
		log.Fatal(err)
	}

	data := struct {
		Issues []shu.GitHubIssue
		Year   int
	}{
		shoes.I,
		time.Now().Year(),
	}

	tmpl.ExecuteTemplate(w, "index", data)
}
