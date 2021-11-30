// Command submit is a chromedp example demonstrating how to fill out and
// submit a form.
package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

func main() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var res string
	var innerHtml string
	err := chromedp.Run(ctx, submit(`https://github.com/search`, `//input[@name="q"]`, `chromedp`, &res, &innerHtml))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("got: `%s\n%s`", strings.TrimSpace(res), innerHtml)
	//read write the html
	html := []byte(innerHtml)
	data := bytes.NewReader(html)

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(data)
	if err != nil {
		log.Fatal(err)
	}
	// Find the review items
	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title, _ := s.Attr("href")
		fmt.Printf("Review %d: %s\n", i, title)
	})

}

func submit(urlstr, sel, q string, res *string, innerHtml *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible(sel),
		chromedp.SendKeys(sel, q),
		chromedp.Submit(sel),
		chromedp.WaitNotPresent(`//*[@id="js-pjax-container"]//h2[contains(., 'Search more than')]`),
		chromedp.Text(`(//*[@id="js-pjax-container"]//ul[contains(@class, "repo-list")]/li[1]//p)[1]`, res),
		chromedp.InnerHTML(`//*[@id="js-pjax-container"]//ul[contains(@class, "repo-list")]`, innerHtml),
	}

}
