package main

import (
    "following-system/routes"
    "net/http"
    "log"
)

func main() {
    router := routes.InitializeRoutes()
    log.Fatal(http.ListenAndServe(":8080", router))
}
