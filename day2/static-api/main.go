package main

import "static/routes"

func main() {

	e := routes.New()

	e.Logger.Fatal(e.Start(":5000"))
}
