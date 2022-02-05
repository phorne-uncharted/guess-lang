package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
	"github.com/pkg/errors"

	log "github.com/unchartedsoftware/plog"
)

func main() {
	pages := crawlWikiFr(1000)
	fmt.Printf("PAGES: %v\n", pages)

	err := output("test.txt", strings.Join(pages, "\n"))
	if err != nil {
		fmt.Printf("error outputing data: %+v", err)
		return
	}
}

func crawlWikiFr(max int) []string {
	log.Infof("pulling links")
	links := []string{}
	c := colly.NewCollector(
		colly.AllowedDomains("fr.wikipedia.org"),
	)

	q, _ := queue.New(1, &queue.InMemoryQueueStorage{MaxSize: 10000})

	// Find and visit all links
	count := 0
	visited := map[string]bool{}
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		//fmt.Printf("CHECKING LINKS: %v\n", e.Attr("href"))
		link := e.Attr("href")
		if strings.Contains(link, "/wiki/") && strings.Index(link, ":") == -1 {
			if !visited[link] {
				visited[link] = true
				//fmt.Printf("HAS WIKI (COUNT: %v): %v\t\n", count, link)
				count++
				if count < max {
					//fmt.Printf("ADDING LINKS: %v\n", link)
					if count%50 == 0 {
						log.Infof("processed %v links", count)
					}
					time.Sleep(500 * time.Millisecond)
					e.Request.Ctx.Put("parent", e.Request.URL.String())
					e.Request.Visit(link)
				}
			}
		}
	})

	c.OnResponse(func(r *colly.Response) {
		links = append(links, r.Request.URL.String())
	})

	q.AddURL("https://fr.wikipedia.org/wiki/(36110)_1999_RV122")
	q.Run(c)

	return links
}

func output(filename string, content string) error {
	err := ioutil.WriteFile(filename, []byte(content), os.ModePerm)
	if err != nil {
		return errors.Wrapf(err, "unable to store content")
	}

	return nil
}
