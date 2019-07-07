package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

//
// Listen to http requests at port 8080 and dump the header and body into a file in the
// directory specified at the command line. If no directory is specifed $PWD is used.
//
// Each request gets dumped into a separate file named YYYY-MM-DD_HH.MM.SS.nnnnnnnnn.hook
//
// The purpose of this is to capture webhook -requests from GitHub for processing by a
// separate process.
//

func main() {
	dir := "."
	if len(os.Args) > 1 {
		dir := os.Args[1]
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Create(dir + "/" + time.Now().Format("2006-01-02_15.04.05.000000000") + ".hook")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		requestDump, _ := httputil.DumpRequest(r, true)
		fmt.Fprint(f, string(requestDump))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
