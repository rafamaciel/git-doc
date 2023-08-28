package server

import (
	"fmt"
	"net/http"
	"path"
	"strings"
)

type Server struct {
	Host string
	Port string

	Source    string
	IndexFile string
}

func (s Server) GetHost() string {
	if s.Host == "" {
		s.Host = "localhost"
	}
	return s.Host
}

func (s Server) GetPort() string {
	if s.Port == "" {
		s.Port = "8080"
	}
	return s.Port
}

func (s Server) GetIndexFile() string {
	if s.IndexFile == "" {
		s.IndexFile = "index.html"
	}
	return s.IndexFile
}

func (s Server) GetSourcePath() string {
	if s.Source == "" {
		s.Source = "docs"
	}
	return s.Source
}

func (s Server) GetAddress() string {

	return strings.Join([]string{s.Host, s.Port}, ":")
}

func (s Server) Start() {

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		filePath := path.Join(s.GetSourcePath(), r.URL.Path)

		if strings.HasSuffix(r.URL.Path, "/") {
			filePath = path.Join(filePath, s.GetIndexFile())
		}

		http.ServeFile(w, r, filePath)
	}))

	sa := s.GetAddress()

	fmt.Printf("server running on http://%s", sa)

	err := http.ListenAndServe(sa, nil)
	if err != nil {
		panic(err)
	}

}
