package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"

	pcbatest "github.com/TheCacophonyProject/pcba-test"
)

func main() {
	if err := runMain(); err != nil {
		log.Fatal(err)
	}
}

var version = "<not set>"

func runMain() error {
	log.SetFlags(0)
	log.Printf("running version: %s", version)

	router := mux.NewRouter()
	static := packr.NewBox("./static")

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/test-attiny", handleTestAttiny)
	apiRouter.HandleFunc("/test-rtc", handleTestRTC)
	apiRouter.HandleFunc("/test-speaker", handleTestSpeakers)
	apiRouter.HandleFunc("/test-usb", handleTestUSB)

	router.Handle("/", http.FileServer(static))

	log.Fatal(http.ListenAndServe(":80", router))
	return nil
}

func handleTestRTC(w http.ResponseWriter, r *http.Request) {
	t := pcbatest.Tests{}
	log.Println("testing rtc")
	pcbatest.TestRTC(1, &t)
	b, err := json.Marshal(t)
	if err != nil {
		serverError(&w, err)
		return
	}
	w.Write(b)
}

func handleTestUSB(w http.ResponseWriter, r *http.Request) {
	t := pcbatest.Tests{}
	log.Println("testing usb")
	pcbatest.TestUSB(1, &t)
	b, err := json.Marshal(t)
	if err != nil {
		serverError(&w, err)
		return
	}
	w.Write(b)
}

func handleTestAttiny(w http.ResponseWriter, r *http.Request) {
	t := pcbatest.Tests{}
	log.Println("testing attiny")
	pcbatest.TestAttiny(&t)
	b, err := json.Marshal(t)
	if err != nil {
		serverError(&w, err)
		return
	}
	w.Write(b)
}

func handleTestSpeakers(w http.ResponseWriter, r *http.Request) {
	t := pcbatest.Tests{}
	log.Println("testing speakers")
	pcbatest.TestSpeakers(&t)
	b, err := json.Marshal(t)
	if err != nil {
		serverError(&w, err)
		return
	}
	w.Write(b)
}

func makeHandler(data string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, data)
	}
}

func serverError(w *http.ResponseWriter, err error) {
	log.Printf("server error: %v", err)
	(*w).WriteHeader(http.StatusInternalServerError)
}
