package controllers

import "github.com/labstack/echo/v4"

func IndexHandler(c echo.Context) error {
	return c.File("static/index.html")
}
