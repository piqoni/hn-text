package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Article struct {
	Title        string
	Link         string
	Comments     int
	CommentsLink string
}

func parseArticles(htmlContent string) ([]Article, error) {
	var articles []Article

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		return nil, err
	}

	doc.Find("tr.athing").Each(func(i int, s *goquery.Selection) {
		title := s.Find("td.title > span.titleline > a").Text()
		link, _ := s.Find("td.title > span.titleline > a").Attr("href")
		commentText := s.Next().Find("a[href^='item']").Last().Text()
		commentsCount, err := extractNumberFromString(commentText)
		commentsLink := s.Next().Find("a[href^='item']").Last().AttrOr("href", "")
		if err != nil {
			commentsCount = 0
		}

		article := Article{
			Title:        title,
			Link:         link,
			Comments:     commentsCount,
			CommentsLink: commentsLink,
		}

		articles = append(articles, article)
	})

	return articles, nil
}

func extractCommentsCount(s *goquery.Selection) (int, error) {
	// find the second a[href^='item'] element

	commentText := s.Next().Find("a[href^='item']").Last().Text()

	return extractNumberFromString(commentText)
}

func extractNumberFromString(input string) (int, error) {

	input = strings.TrimSpace(input)
	re := regexp.MustCompile(`\d+`)
	matches := re.FindString(input)
	if matches == "" {
		return 0, fmt.Errorf("no numbers found in input")
	}

	number, err := strconv.Atoi(matches)
	if err != nil {
		return 0, err
	}

	return number, nil
}
