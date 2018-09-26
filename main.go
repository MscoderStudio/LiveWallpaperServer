package main

import (
	"net/http"
	"LiveWallpaperServer/upupoo"
	"time"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

var config Configuration

func index(writer http.ResponseWriter, request *http.Request) {

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
