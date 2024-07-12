package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/test/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, c.Param("id"))
	})

	e.GET("/health", helthcheckHandler)
	e.GET("/web", helthcheckHandler)

	nestedGroup := e.Group("/nested")
	{
		nestedGroup.GET("", nestedCheckerHandler)
		nestedNestedGroup := nestedGroup.Group("/2deps")
		{
			nestedNestedGroup.GET("", nestedCheckerHandler)
		}
	}

	e.Logger.Fatal(e.Start(":1323"))
}

func helthcheckHandler(c echo.Context) error {
	return c.HTML(http.StatusOK, "<h1>빠이</h1>")
}

func nestedCheckerHandler(c echo.Context) error {
	return c.HTML(http.StatusOK, "<h1>네스티드</h1>")
}
