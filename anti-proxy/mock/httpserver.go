package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type webServer1 struct{}

func (ws1 webServer1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello server1")
}

type webServer2 struct{}

func (ws2 webServer2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello server2")
}

func main() {
	c := make(chan os.Signal)

	go func() {
		http.ListenAndServe(":9090", webServer1{})
	}()

	go func() {
		http.ListenAndServe(":9091", webServer2{})
	}()

	signal.Notify(c, os.Interrupt)

	s := <-c
	log.Println(s)
}
