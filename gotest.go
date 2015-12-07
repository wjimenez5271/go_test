package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
        fmt.Printf("starting the program\n")
        //log.Print("starting the program")
	res, err := http.Get("http://www.willjimenez.com/")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

