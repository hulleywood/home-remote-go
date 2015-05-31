package main

import (
    "net/http"
)

func LivingRoomMonitor(w http.ResponseWriter, r *http.Request) {

    SetDefaultHeaders(w)

    w.WriteHeader(200)
}

func BedRoomMonitor(w http.ResponseWriter, r *http.Request) {

    SetDefaultHeaders(w)

    w.WriteHeader(200)
}

func SendAction(w http.ResponseWriter, r *http.Request) {

    SetDefaultHeaders(w)

    w.WriteHeader(200)
}

func SetDefaultHeaders(w http.ResponseWriter) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.Header().Set("Server", "Home Remote Go Server")
}
