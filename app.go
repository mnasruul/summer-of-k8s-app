package main

import (
	"fmt"
	"net/http"
)

func main() {

	version := "<b>1.3.0</b>"

	color := "#c034eb" //Blue 44B3C2 and Yellow F1A94E

	http.HandleFunc("/callme", func(w http.ResponseWriter, r *http.Request) {
		headers := ""
		for i, v := range r.Header {
			headers = headers + fmt.Sprintf("<b>%s</b>: %s\n", i, v)
			// fmt.Fprintf(w, "<div class='pod' style='background:%s'> %s: %s\n </div>", color, i, v)
		}
		fmt.Fprintf(w, "<div class='pod' style='background:%s'> ver: %s\n headers:\n %s </div>", color, version, headers)
	})

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("Listening now at port 8080")
	http.ListenAndServe(":8080", nil)
}
