package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	//Add a GPL3 package to cause havock
	// os.Setenv("test", string(exitcodes.Success))

	c := os.Getenv("COLOR")
	if len(c) == 0 {
		os.Setenv("COLOR", "#44B3C2") //Blue 44B3C2 and Yellow F1A94E
	}

	http.HandleFunc("/callme", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<html style='background:"+os.Getenv("COLOR")+"'> Requested: %s\n </html>", r.URL.Path)
	})

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("Listening now at port 8080")
	http.ListenAndServe(":8080", nil)
}
