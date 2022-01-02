package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	for i := 1; i <= 6; i++ {

		fmt.Fprintf(w, "<h%v>Hello world!</h%v>", i)

	}

}

func main() {

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)

}
