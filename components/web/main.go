package main

import "net/http"

func main() {
	fs := http.FileServer(http.Dir("components/web/static"))

	http.Handle("/", fs)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
