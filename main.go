package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

var dir=""

func main() {
	dir=os.Args[1]

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Create(dir+time.Now().Format("2006-01-02_15.04.05_000000000")+".hook")
		defer f.Close()
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Fprint(f, err.Error())
		} else {
			fmt.Fprint(f, string(requestDump))
		}
		fmt.Fprintf(w,"OK")
	})
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}


