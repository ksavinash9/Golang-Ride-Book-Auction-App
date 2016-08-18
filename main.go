package main


import (

    "./routers"
    "./models"
    "./conf"
    "fmt"
    "log"
    "net/http"
)


func main() {
    // // Configuration
    PORT := config.Get("WEBSERVER_PORT")
    PG_DATABASE_NAME := config.Get("DATABASE_NAME")
    PG_DATABASE_USER := config.Get("DATABASE_USER")
    PG_SSL_MODE := config.Get("SSL_MODE")

    models.InitDB("user=" + PG_DATABASE_USER + " dbname=" + PG_DATABASE_NAME + " sslmode=" + PG_SSL_MODE)

    // Get the router from router.go
    router := routers.GetRouter()

    //Run HTTP Server
    fmt.Println("Running WebServer on Port " + PORT)

    log.Fatal(http.ListenAndServe(":" + PORT, router))
}