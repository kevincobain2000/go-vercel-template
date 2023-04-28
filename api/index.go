package main

// package handler

import (
	"net/http"

	app "github.com/kevincobain2000/go-vercel-template/app"
)

// comment this out to use the handler package
func main() {
	e := app.HTTPServer()
	e.Logger.Fatal(e.Start("localhost:3000"))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	e := app.HTTPServer()
	e.ServeHTTP(w, r)
}
