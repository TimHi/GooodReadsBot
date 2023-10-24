package main

import (
	"net/http"
	"net/url"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"github.com/timhi/discordgoodreads/cmd"
)

func main() {
	e := echo.New()
	e.GET("/book/:query", func(c echo.Context) error {
		query := c.Param("query")
		book, err := cmd.SearchBook(url.PathEscape(query))
		if err != nil {
			log.Error(err)
		}
		return c.JSON(http.StatusOK, book)
	})
	e.Logger.Fatal(e.Start(":7000"))
}
