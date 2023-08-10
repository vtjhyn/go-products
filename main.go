package main

import (
	"belajar/config"
	"belajar/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {
config.ConnectDB()
http.HandleFunc("/", homecontroller.Welcome)
log.Println("Server running on port 8080") 
http.ListenAndServe(":8080",nil)
}