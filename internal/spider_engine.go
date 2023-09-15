package internal

import (
	"fmt"
	"log"
	"regexp"
	"spider/internal/pages"

	"github.com/gocolly/colly/v2"
)

type SpiderEngine struct {
	pages []pages.BasePage
}

func NewSpiderEngine() *SpiderEngine {
	wtaPage := pages.NewWTATourPage()
	pages := []pages.BasePage{
		wtaPage,
	}

	return &SpiderEngine{pages: pages}
}

func (engine *SpiderEngine) Start() {
	depthFilter := colly.MaxDepth(2)
	ignoreRobotsFilter := colly.IgnoreRobotsTxt()
	userAgent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.5.2 Safari/605.1.15"
	userAgentFilter := colly.UserAgent(userAgent)
	proxyURL := "http://192.168.50.33:8888"

	for _, page := range engine.pages {
		urlFilter := colly.URLFilters(regexp.MustCompile(page.FilterURL()))
		homeCollector := colly.NewCollector(userAgentFilter, ignoreRobotsFilter, urlFilter, depthFilter)
		homeCollector.SetProxy(proxyURL)

		// Find and visit all links
		homeCollector.OnHTML("a[href]", func(e *colly.HTMLElement) {
			page.HandleLink(homeCollector, e)
		})

		homeCollector.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL)
			r.Headers.Set("Accept", "*/*")
		})

		homeCollector.OnError(func(r *colly.Response, err error) {
			log.Fatal("error")
		})

		homeCollector.Visit(page.HomeURL())
		homeCollector.Wait()
	}
}
