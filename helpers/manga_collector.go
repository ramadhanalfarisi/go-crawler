package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
	"github.com/ramadhanalfarisi/go-crawler/model"
)

func CrawlManga(mg *[]model.Manga) {
	c := colly.NewCollector(
		colly.AllowedDomains("myanimelist.net", "www.myanimelist.net"),
		colly.MaxDepth(1),
	)
	infoCollector := c.Clone()

	c.OnHTML(".ranking-list", func(e *colly.HTMLElement) {
		profileUrl := e.ChildAttr("td.title > a", "href")
		profileUrl = e.Request.AbsoluteURL(profileUrl)
		infoCollector.Visit(profileUrl)
	})

	c.OnHTML("a.next", func(e *colly.HTMLElement) {
		nextPage := e.Request.AbsoluteURL(e.Attr("href"))
		c.Visit(nextPage)
	})

	infoCollector.OnHTML("#contentWrapper", func(e *colly.HTMLElement) {
		manga := model.Manga{}
		statistic := model.Statistic{}
		manga.Title = e.ChildText("span.h1-title > span")
		statistic.Score = e.ChildText("span.score-label")
		statistic.Ranked = e.ChildText("span.ranked > strong")
		statistic.Popularity = e.ChildText("span.popularity > strong")
		manga.Volumes = strings.TrimSpace(e.ChildText("#totalVols"))
		manga.Chapters = strings.TrimSpace(e.ChildText("#totalChaps"))
		manga.Authors = strings.TrimSpace(e.ChildText("span.author"))

		e.ForEach("div.spaceit_pad", func(i int, h *colly.HTMLElement) {
			label := h.ChildText("span.dark_text")
			switch label {
			case "Published:":
				manga.Year = strings.ReplaceAll(strings.TrimSpace(h.Text), "Published: ", "")
			case "Genres:":
				h.ForEach("a", func(i int, j *colly.HTMLElement) {
					if i != 0 {
						manga.Genres += ","
					}
					manga.Genres += j.Attr("title")
				})
			case "Genre:":
				h.ForEach("a", func(i int, j *colly.HTMLElement) {
					if i != 0 {
						manga.Genres += ","
					}
					manga.Genres += j.Attr("title")
				})
			case "Themes:":
				h.ForEach("a", func(i int, j *colly.HTMLElement) {
					if i != 0 {
						manga.Themes += ","
					}
					manga.Themes += j.Attr("title")
				})
			case "Theme:":
				h.ForEach("a", func(i int, j *colly.HTMLElement) {
					if i != 0 {
						manga.Themes += ","
					}
					manga.Themes += j.Attr("title")
				})
			}

		})

		manga.Statistic = statistic
		*mg = append(*mg, manga)
		json, err := json.Marshal(manga)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(string(json))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL.String())
	})
	infoCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting Profile URL: ", r.URL.String())
	})

	c.Visit("https://myanimelist.net/topmanga.php")
}
