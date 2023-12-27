package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// trays:6009/modpacks
// trays:6009/resources/:modpackname/:files
// trays:6009/files/:modpackname/:files
func routes(e *echo.Echo) {
	e.GET("modpacks", func(c echo.Context) error { return c.JSON(http.StatusOK, Modpacks) })
	e.GET("resources/:modpack/:file", M_Resources)
	e.GET("files/:modpack/:file", M_Files)
}

func M_Resources(c echo.Context) error {
	modpack := c.Param("modpack")
	file := c.Param("file")
	return c.File(TRAYS_FOLDER + "storage/public/resources/" + modpack + "/" + file)
}

func M_Files(c echo.Context) error {
	modpack := c.Param("modpack")
	file := c.Param("file")
	return c.File(TRAYS_FOLDER + "storage/public/modpacks/" + modpack + "/" + file)
}
