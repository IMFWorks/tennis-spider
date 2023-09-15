package pages

import "github.com/gocolly/colly/v2"

type BasePage interface {
	HomeURL() string
	FilterURL() string
	HandleLink(homeCollector *colly.Collector, htmlElement *colly.HTMLElement)
}
