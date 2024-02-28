package main

import (
	"appstore/backend"
	"appstore/handler"
	"fmt"
	"log"
	"net/http"
)

/*
type MyHandler struct {}

	func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	    // Handle incoming requests here
		fmt.Println("Eason received one upload request2")
		decoder := json.NewDecoder(r.Body)
		var app model.App
		if err := decoder.Decode(&app); err != nil{
			panic(err)
		}
		fmt.Fprintf(w, "Upload request have been received2: %s\n", app.Description)
	}
*/
func main() {
	fmt.Println("Eason, my master")
	fmt.Println("started-service")

	//m := MyHandler{}
	//http.HandleFunc("/eason", m.ServeHTTP)
	backend.InitElasticsearchBackend()
	backend.InitGCSBackend()

	log.Fatal(http.ListenAndServe(":8080", handler.InitRouter()))
	//log.Fatal(http.ListenAndServe(":8080", m))
}
