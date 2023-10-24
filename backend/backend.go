package backend

import (
	"net/http"
	"net/url"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"github.com/timhi/GooodReadsBot/backend/cmd"
)

func Start(port string, ch chan<- bool) {
	e := echo.New()
	e.GET("/book/:query", func(c echo.Context) error {
		query := c.Param("query")
		book, err := cmd.SearchBook(url.PathEscape(query))
		if err != nil {
			log.Error(err)
		}
		return c.JSON(http.StatusOK, book)
	})
	e.Logger.Fatal(e.Start(port))

}
