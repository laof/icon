package tools

import (
	"context"
	"log"
	"path"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func GetLink(url string, newPath chan string) {

	// 禁用chrome headless
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 25*time.Second)
	defer cancel()

	var res []map[string]string
	var base string = ""
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitEnabled(`head > link`, chromedp.ByQuery),
		chromedp.Sleep(1*time.Second),
		chromedp.AttributesAll("head > link", &res),
		chromedp.Evaluate("document.baseURI", &base),
	)
	if err != nil {
		newPath <- err.Error()
		return
	}

	// var links = []string{}

	for _, k := range res {
		href := any(k["href"], base)
		if href != "" {
			// links = append(links, href)
			newPath <- href
			return
		}
	}

	newPath <- ""
}

func any(href, base string) string {
	var types = []string{".ico", ".png", ".svg", ".gif"}
	for i := range types {
		if strings.HasSuffix(href, types[i]) {
			return link(href, base)
		}
	}
	return ""
}

func link(href, base string) string {
	if strings.HasPrefix(href, "//") {
		return "https:" + href
	} else if strings.HasPrefix(href, "http") {
		return href
	}
	return path.Join(base, href)
}
