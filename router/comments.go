package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func PostComments(c echo.Context) error {
	return c.String(http.StatusOK, "PostComments")
}

func PutComments(c echo.Context) error {
	return c.String(http.StatusOK, "PutComments")
}

func DeleteComments(c echo.Context) error {
	return c.String(http.StatusOK, "DeleteComments")
}
