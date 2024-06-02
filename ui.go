package main

import (
	"fmt"
	"log"
	"net/url"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/gelembjuk/articletext"
	"github.com/rivo/tview"
)

func createArticleList(articles []Article) *tview.List {
	list := tview.NewList().ShowSecondaryText(true).SetSecondaryTextColor(tcell.ColorGray)
	for _, article := range articles {
		title := article.Title
		if article.Comments > 50 { // TODO: configurable
			title = "🔥 " + title
		}

		commentsText := strconv.Itoa(article.Comments) + " comments"
		list.AddItem(title, extractDomain(article.Link)+" "+commentsText, 0, nil)
	}

	return list
}

func extractDomain(link string) string {
	u, err := url.Parse(link)
	if err != nil {
		log.Fatal(err)
	}
	return u.Host
}

func createInputHandler(app *tview.Application, list *tview.List, articles []Article, pages *tview.Pages) func(event *tcell.EventKey) *tcell.EventKey {

	return func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlC:
			app.Stop()
			return nil
		case tcell.KeyRight:
			nextPage(pages, app, articles, list)
			return nil
		case tcell.KeyLeft:
			backPage(pages)
			return nil
		case tcell.KeyRune:
			switch event.Rune() {
			case 'q':
				app.Stop()
				return nil
			case 'j':
				return tcell.NewEventKey(tcell.KeyDown, 0, tcell.ModNone)
			case 'k':
				return tcell.NewEventKey(tcell.KeyUp, 0, tcell.ModNone)
			case 'l':
				nextPage(pages, app, articles, list)
				return nil
			case 'h':
				backPage(pages)
				return nil
			case ' ':
				openURL(articles[list.GetCurrentItem()].Link)
				return nil
			case 'c':
				openURL(hackerNewsURL + articles[list.GetCurrentItem()].CommentsLink)
				return nil

			}
		}

		return event
	}
}

func backPage(pages *tview.Pages) {
	// TODO: navigation flow will become configurable
	currentPage, _ := pages.GetFrontPage()
	if currentPage == "comments" {
		pages.SwitchToPage("homepage")
	}
	if currentPage == "article" {
		pages.SwitchToPage("comments")
	}
}

func nextPage(pages *tview.Pages, app *tview.Application, articles []Article, list *tview.List) {
	currentPage, _ := pages.GetFrontPage()
	if currentPage == "comments" {
		openArticle(app, articles[list.GetCurrentItem()].Link, pages)
	} else {
		openComments(app, articles[list.GetCurrentItem()].CommentsLink, pages)
	}
}

func openComments(app *tview.Application, commentsLink string, pages *tview.Pages) {
	u, err := url.Parse(commentsLink)
	if err != nil {
		fmt.Println("Error parsing URL:", err) // TODO maybe alert dialogbox
		return
	}
	story_id := u.Query().Get("id")

	articleStringList := fetchComments(story_id)
	commentsText := ""
	for _, articleString := range articleStringList {
		commentsText += articleString + "\n"
	}
	displayComments(app, pages, commentsText)
}

func openArticle(app *tview.Application, articleLink string, pages *tview.Pages) {
	articleText := getArticleTextFromLink(articleLink)
	displayArticle(app, pages, articleText)
}

func getArticleTextFromLink(url string) string {
	article, err := articletext.GetArticleTextFromUrl(url)
	if err != nil {
		fmt.Printf("Failed to parse %s, %v\n", url, err)
	}
	return article
}

func displayArticle(app *tview.Application, pages *tview.Pages, text string) {
	articleTextView := tview.NewTextView().
		SetText(text).
		SetDynamicColors(true).
		SetScrollable(true)

	pages.AddPage("article", articleTextView, true, true)
	if err := app.SetRoot(pages, true).Run(); err != nil {
		log.Fatal(err)
	}
}

func displayComments(app *tview.Application, pages *tview.Pages, text string) {
	commentsTextView := tview.NewTextView().
		SetText(text).
		SetDynamicColors(true).
		SetScrollable(true)

	pages.AddPage("comments", commentsTextView, true, true)
	if err := app.SetRoot(pages, true).Run(); err != nil {
		log.Fatal(err)
	}
}

func openURL(url string) {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	exec.Command(cmd, args...).Start()
}
