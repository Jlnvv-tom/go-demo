package main

import (
	"fmt"
	"net/http"
)

func morning(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "morning!")
}

func evening(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "evening!")
}

type Morning struct{}

func (m Morning) Start() {
	http.HandleFunc("/morning", morning)
	http.ListenAndServe("localhost:8080", nil)
}

func (m Morning) Stop() {
	fmt.Println("Stop morning service...")
}

type Evening struct{}

func (e Evening) Start() {
	http.HandleFunc("/evening", evening)
	http.ListenAndServe("localhost:8081", nil)
}

func (e Evening) Stop() {
	fmt.Println("Stop evening service...")
}

func main() {
	var morning Morning
	morning.Start()
	defer morning.Stop()

	var evening Evening
	evening.Start()
	defer evening.Stop()

}
