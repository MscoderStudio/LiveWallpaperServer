package main

import (
	"LiveWallpaperServer/upupoo"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

var config Configuration

func getTags(writer http.ResponseWriter, request *http.Request) {
	tags, err := upupoo.GetTags()
	b, err := json.Marshal(tags.Data)
	if err != nil {
		log.Fatal(err)
	}
	// result = string(b[:])
	writer.Write(b)
}

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}

func main() {
	loadConfig()
	mux := http.NewServeMux()
	mux.HandleFunc("/tags", getTags)

	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
