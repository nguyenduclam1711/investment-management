package main

import (
	"log"

	"github.com/nguyenduclam1711/investment-management/app"
)

func main() {
	app := app.App()

	log.Fatal(app.Listen(":3000"))
}
