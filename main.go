package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/yuichi10/pcf-app/util"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
	fmt.Fprintln(w, "5 + 3=", util.Add(5, 3))
	fmt.Fprintln(w, "5 - 3=", util.Sub(5, 3))
}

func main() {
	port := os.Getenv("PORT")
	if len(port) <= 0 {
		port = "8888"
	}
	http.HandleFunc("/", hello)
	http.ListenAndServe(":"+port, nil)
}
