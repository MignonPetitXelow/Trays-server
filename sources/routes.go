package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// trays/get/modpacks
// trays/get/:modpackname
// trays/get/:modpackname/:files
func routes(e *echo.Echo) {
	e.GET("modpacks", func(c echo.Context) error { return c.JSON(http.StatusOK, Modpacks) })
	e.GET("resources/:modpack/:file", func(c echo.Context) error { return Resources(c) })
}

func Resources(c echo.Context) error {
	return c.File(TRAYS_FOLDER + "storage/public/resources/" + c.Param("modpack") + "/" + c.Param("file"))
}
