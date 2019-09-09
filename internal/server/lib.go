package server

import (
	"fmt"
	"html/template"
	"os"
	"os/signal"
	"syscall"
)

// parseMany is like ParseFiles, but instead of files takes
// an argument of `map[name]content`, where name and content are both strings.
// Based on html/template/parseFiles() from the standard library.
func parseMany(args map[string]string) (*template.Template, error) {
	var t *template.Template
	for name, content := range args {
		var tmpl *template.Template
		if t == nil {
			t = template.New(name)
		}

		if name == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(name)
		}

		_, err := tmpl.Parse(content)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

// Thanks Edd: https://golangcode.com/handle-ctrl-c-exit-in-terminal/
func setupCloseHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\rCtrl+C pressed in Terminal")
		os.Exit(0)
	}()
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}
