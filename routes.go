package main

import (
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "LivingRoomMonitor",
        "GET",
        "/living_room/monitor",
        LivingRoomMonitor,
    },
    Route{
        "BedRoomMonitor",
        "GET",
        "/bed_room/monitor",
        BedRoomMonitor,
    },
    Route{
        "SendAction",
        "GET",
        "/send/{device}/{action}",
        SendAction,
    },
}
