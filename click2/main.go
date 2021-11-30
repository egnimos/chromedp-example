// Command click2 is a chromedp example demonstrating how to use a selector to
// click on an element.
package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

func main() {
	scrap()
}

// // create chrome instance
// ctx, cancel := chromedp.NewContext(
// 	context.Background(),
// 	// chromedp.WithDebugf(log.Printf),
// )
// defer cancel()

// // create a timeout
// ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
// defer cancel()

// // navigate to a page, wait for an element, click
// var example string
// var output string
// err := chromedp.Run(ctx,
// 	chromedp.Navigate(`https://pkg.go.dev/time`),
// 	// wait for footer element is visible (ie, page is loaded)
// 	chromedp.WaitVisible(`body > footer`),
// 	// find and click "Example" link
// 	chromedp.Click(`#example-After`, chromedp.NodeVisible),
// 	// retrieve the text of the textarea
// 	chromedp.Value(`#example-After textarea`, &example),

// 	//click the run button
// 	chromedp.Click(`#example-After button.Documentation-exampleRunButton`, chromedp.ByQuery),
// 	// chromedp.Button()()
// 	// //retrieve the text from the output
// 	chromedp.Sleep(5*time.Second),
// 	chromedp.Text(`#example-After span.Documentation-exampleOutput`, &output, chromedp.ByQueryAll),
// )
// if err != nil {
// 	log.Fatal(err)
// }
// log.Printf("Go's time.After example:\n%s\n%s", example, output)

func scrap() {
	// https://www2.kickassanime.ro/anime/gintama-2015--449963
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()
	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	// navigate to a page, wait for an element, click
	var example string
	var innerHtml string
	// var attribute string
	var attributes map[string]string
	var outerHtml string
	tasks := chromedp.Tasks{
		chromedp.Navigate("https://www2.kickassanime.ro"),
		chromedp.WaitVisible(`footer#footer`),
		chromedp.InnerHTML(`#main-video-list div.video-list.row.mx-0`, &innerHtml),
	}
	err := chromedp.Run(ctx,
		// wait for footer element is visible (ie, page is loaded)
		// chromedp.WaitVisible(`footer#footer`),
		// //chrome dp
		// chromedp.TextContent(`#content p.mb-0`, &example),
		// chromedp.Attributes(`a[href]`, &attributes),
		// // chromedp.Attribute(`a[href]`, &attribute),
		// chromedp.InnerHTML(`#content`, &innerHtml),
		// chromedp.OuterHTML(`#content`, &outerHtml),
		tasks,
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Go's time.After example:\n%s\n%s\n%s", example, attributes, outerHtml)
	log.Printf("INNERHTML:%s", innerHtml)
	byteData := []byte(innerHtml)
	readerData := bytes.NewReader(byteData)

	//process the html and parse it to the main json value
	doc, err := goquery.NewDocumentFromReader(readerData)
	if err != nil {
		panic(err)
	}

	doc.Find("div.video-item.col-6.mb-2.px-1").Each(func(i int, s *goquery.Selection) {
		navigationUrl := s.Find("a.ka-url-wrapper.video-item-poster.rounded").AttrOr("href", "empty")
		imageUrl := s.Find("a.ka-url-wrapper.video-item-poster.rounded").AttrOr("style", "empty")
		parsedUrl := bytes.Split([]byte(imageUrl), []byte("\""))
		title := s.Find("a.ka-url-wrapper.video-item-title").Text()
		html, _ := s.Html()
		fmt.Printf("%s\n%s\n%s\n%s", navigationUrl, parsedUrl[1], title, html)
	})
}
