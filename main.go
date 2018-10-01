package main

import (
	"SKSISAPI/router"
	"log"
	"net/http"
)

func main()  {
	router.Routing()

	err := http.ListenAndServe(":9006",nil)
	if err !=nil{
		log.Fatal(err)
	}
}