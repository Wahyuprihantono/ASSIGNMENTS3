package main

import (
	"ASSIGNMENTS3/static"
	"fmt"
	"net/http"
)

func main() {
	go static.AutoReload()
	http.HandleFunc("/", static.ReloadWeb)
	fmt.Println("server running on PORT:", ":8080")
	http.ListenAndServe(":8080", nil)
}
