package main

import (
	"encoding/json"
	"flag"
	"net/http"
	"net/url"
)

type Details struct {
	RemoteAddr string
	Host       string
	Method     string
	URL        string
	Header     http.Header
	Form       url.Values
}

func getDetails(r *http.Request) *Details {
	r.ParseForm()
	return &Details{
		RemoteAddr: r.RemoteAddr,
		Host:       r.Host,
		Method:     r.Method,
		URL:        r.URL.String(),
		Header:     r.Header,
		Form:       r.Form,
	}
}

func main() {
	addr := flag.String("addr", ":8080", "HTTP listen address")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bytes, err := json.MarshalIndent(getDetails(r), "", "    ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(bytes)
	})
	http.ListenAndServe(*addr, nil)
}
