package pages

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

type WTATourPage struct {
	homeURL   string
	filterURL string
	source    string
}

func NewWTATourPage() *WTATourPage {
	return &WTATourPage{
		homeURL:   "https://www.wtatennis.com",
		filterURL: "https://www.wtatennis.com",
		source:    "WTA",
	}
}

func (page *WTATourPage) HomeURL() string {
	return page.filterURL
}

func (page *WTATourPage) FilterURL() string {
	return page.filterURL
}

func (page *WTATourPage) HandleLink(homeCollector *colly.Collector, htmlElement *colly.HTMLElement) {
	href := htmlElement.Attr("href")
	className := htmlElement.Attr("class")
	if className == "content-listing-item__link" {
		articleCollector := homeCollector.Clone()

		onArticle(articleCollector)

		articleURL := htmlElement.Request.AbsoluteURL(href)
		articleCollector.Visit(articleURL)
	}
}

func onArticle(articleCollector *colly.Collector) {
	// Find Article content
	articleCollector.OnHTML("#main-content > article", func(e *colly.HTMLElement) {
		sourceURL := e.Request.URL.String()
		fmt.Println("sourceURL:", sourceURL)
		title := e.ChildText("header > div.content-page-header__info > h1")
		fmt.Println("title:", title)

		// get thumb image
		thumbUrl := e.ChildAttr("source", "srcset")
		fmt.Println("thumb url:", thumbUrl)

		//src="https://photoresources.wtatennis.com/photo-resources/2023/09/11/6bd4da58-fc6c-48be-b8c3-bffb6daa5dea/Coco_Gauff_Jessica_Pegula_-_2023_US_Open_-_Day_8-DSC_1140.jpg?width=1440"

		// content := e.ChildTexts("div.article__body.article-body.wrapper > p")
		// fmt.Println("content:", content)
	})

	articleCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	articleCollector.OnError(func(r *colly.Response, err error) {
		log.Fatal("error")
	})
}
