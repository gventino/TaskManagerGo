package main

import (
	"fmt"
	"main/views"
	"main/views/layout"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	login := layout.Base(views.LoginPage())

	http.Handle("/", templ.Handler(login))

	fmt.Println("Server on at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
