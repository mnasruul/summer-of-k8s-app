package main

import (
	"fmt"
	"net/http"
)

func main() {

	version := "<b>1.3.0</b>"

	color := "#F1A94E" //Blue 44B3C2 and Yellow F1A94E

	http.HandleFunc("/callme", func(w http.ResponseWriter, r *http.Request) {
		headers := ""
		headers = r.Header.Get("mnasruul-summer-of-k8s-app")
		fmt.Fprintf(w, "<div class='pod' style='background:%s'> ver: %s\n headers:\n %s </div>", color, version, headers)
	})

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("Listening now at port 8080")
	http.ListenAndServe(":8080", nil)
}
