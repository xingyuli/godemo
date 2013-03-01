package main

import (
	"fmt"
	"os"
	"net/http"
)

// Counter server.
type Counter int

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	*ctr++
	fmt.Fprintf(w, "counter = %d\n", *ctr)
}

// Argument server.
func FirstApproach() {
	ArgServer := func(w http.ResponseWriter, req *http.Request) {
		for _, s := range os.Args {
			fmt.Fprintln(w, s)
		}
	}
	http.Handle("/args", http.HandlerFunc(ArgServer))
}

func AnotherApproach() {
	http.HandleFunc("/args", func(w http.ResponseWriter, req *http.Request) {
		for _, s := range os.Args {
			fmt.Fprintln(w, s)
		}
	})
}

func main() {
	http.Handle("/counter", new(Counter))
	
	// FirstApproach()
	AnotherApproach()

	http.ListenAndServe("localhost:4000", nil)
}

/* ************************************************************************* */

// A channel that sends a notification on each visit.
// (Probably want the channel to be buffered.)
type Chan chan *http.Request

func (ch *Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	*ch <- req
	fmt.Fprint(w, "notification sent")
}
