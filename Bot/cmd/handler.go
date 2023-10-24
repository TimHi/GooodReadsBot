package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/gocolly/colly"
	"github.com/timhi/GooodReadsBot/Bot/model"
)

func SearchBook(title string) (model.Book, error) {
	var err error = nil
	book := model.Book{}
	c := colly.NewCollector()

	c.OnHTML("table[class='tableList'] > tbody > tr:first-child", func(h *colly.HTMLElement) {
		book.Link = h.ChildAttr("a[class='bookTitle']", "href")
		book.Title = h.ChildText("a.bookTitle > span[itemprop='name']")

		h.ForEach("span[itemprop='author'] a.authorName > span[itemprop='name']", func(_ int, authorElement *colly.HTMLElement) {
			authorName := authorElement.Text
			book.Authors = append(book.Authors, authorName)
		})

		book.Rating = h.ChildText("span[class='minirating']")
	})

	c.OnRequest(func(r *colly.Request) {
		log.Info("Visiting", "url", r.URL)
	})

	c.OnError(func(r *colly.Response, scrapeError error) {
		log.Error("Scraper fucked up", "err", scrapeError)
		err = scrapeError
	})

	c.Visit("https://www.goodreads.com/search?utf8=✓&query=" + title)
	return book, err
}
