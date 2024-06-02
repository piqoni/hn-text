package main

import (
	"log"
	"os"

	"github.com/rivo/tview"
)

var hackerNewsURL = "https://news.ycombinator.com/"

func main() {
	app := tview.NewApplication()
	// TODO: rewrite this for other options
	if len(os.Args) > 1 && os.Args[1] == "best" {
		hackerNewsURL = "https://news.ycombinator.com/best"
	}

	htmlContent, err := fetchWebpage(hackerNewsURL)
	if err != nil {
		log.Fatal(err)
	}

	articles, err := parseArticles(htmlContent)
	if err != nil {
		log.Fatal(err)
	}

	list := createArticleList(articles)
	pages := tview.NewPages()
	pages.AddPage("homepage", list, true, false)

	app.SetInputCapture(createInputHandler(app, list, articles, pages))

	if err := app.SetRoot(list, true).Run(); err != nil {
		log.Fatal(err)
	}
}
