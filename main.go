package main

import "github.com/ramadhanalfarisi/go-crawler/app"

func main() {
	var app app.App
	app.CreateRoutes()
	app.Run()
}