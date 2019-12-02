package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gobuffalo/packr"
	"github.com/godbus/dbus"
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
	apiRouter.HandleFunc("/test-attiny", handleTestAttiny).Methods("GET")
	apiRouter.HandleFunc("/test-rtc", handleTestRTC).Methods("GET")
	apiRouter.HandleFunc("/test-speaker", handleTestSpeakers).Methods("GET")
	apiRouter.HandleFunc("/test-usb", handleTestUSB).Methods("GET")
	apiRouter.HandleFunc("/camera/snapshot", TakeSnapshot).Methods("PUT")

	router.HandleFunc("/camera/snapshot", CameraSnapshot)
	router.PathPrefix("/").Handler(http.FileServer(static))

	log.Fatal(http.ListenAndServe(":80", router))
	return nil
}

func CameraSnapshot(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/var/spool/cptv/still.png")
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

// TakeSnapshot will request a new snapshot to be taken by thermal-recorder
func TakeSnapshot(w http.ResponseWriter, r *http.Request) {
	conn, err := dbus.SystemBus()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	recorder := conn.Object("org.cacophony.thermalrecorder", "/org/cacophony/thermalrecorder")
	err = recorder.Call("org.cacophony.thermalrecorder.TakeSnapshot", 0).Err
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func handleTestUSB(w http.ResponseWriter, r *http.Request) {
	t := pcbatest.Tests{}
	log.Println("testing usb")
	pcbatest.TestUSB(3, &t)
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

func serverError(w *http.ResponseWriter, err error) {
	log.Printf("server error: %v", err)
	(*w).WriteHeader(http.StatusInternalServerError)
}
