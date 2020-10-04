package main

import (
	"./uri"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	router := uri.NewRouter()

	var port string
	if len(os.Args) > 1 {
		port = os.Args[1]
		doMining()
	} else {
		port = "6689"
	}
	fmt.Println("port is: ", port)
	uri.InitSelfAddress(port)

	log.Fatal(http.ListenAndServe(":"+port, router))

}

func doMining (){
	uri.Download()
}