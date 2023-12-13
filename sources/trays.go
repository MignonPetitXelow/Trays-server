package main

import (
	"github.com/labstack/echo/v4"
)

var Modpacks []Pack
var TRAYS_FOLDER string = "./"

func main() {
	e := echo.New()

	Setup_Start()

	routes(e)
	e.Start(":6009")
}