package main

import (
	"fmt"
	"net/http"
)

func setText(msg string) string {
	return "<b>" + msg + "</b>"
}

func showIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, setText("Code.education Rocks!"))
}

func main()  {
	http.HandleFunc("/", showIndex)
	http.ListenAndServe(":8000", nil)
}