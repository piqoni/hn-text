package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/k3a/html2text"
)

func fetchWebpage(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

const HN_SEARCH_URL = "https://hn.algolia.com/api/v1/"

type Comment struct {
	Author   string    `json:"author"`
	Text     string    `json:"text"`
	Children []Comment `json:"children"`
}

func sanitize(input string) string {
	return html2text.HTML2Text(input)
}

func safeRequest(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to fetch URL: %s\n", url)
		return nil
	}
	return resp
}

func fetchComments(storyID string) []string {
	resp := safeRequest(HN_SEARCH_URL + "items/" + storyID)
	if resp == nil {
		return nil
	}
	defer resp.Body.Close()

	var comments map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&comments); err != nil {
		fmt.Printf("Failed to decode JSON response\n")
		return nil
	}

	var lines []string
	lines = append(lines, " ")

	if children, ok := comments["children"].([]interface{}); ok {
		for _, child := range children {
			childComment := child.(map[string]interface{})
			appendComment(childComment, &lines, 0)
		}
	}

	return lines
}

func appendComment(comment map[string]interface{}, lines *[]string, level int) {
	indent := "" + strings.Repeat("   ", min(level, 4)*2) + "| "

	if author, ok := comment["author"].(string); ok {
		*lines = append(*lines, indent+sanitize(author)+" wrote:")

		text := sanitize(comment["text"].(string))

		paragraphs := strings.Split(text, "\n\n")
		for _, paragraph := range paragraphs {
			textLines := wrapText(paragraph, indent)
			*lines = append(*lines, textLines...)
			*lines = append(*lines, indent)
		}
		*lines = (*lines)[:len(*lines)-1] // Drop the blank line after the last paragraph
	} else {
		*lines = append(*lines, indent+"[deleted]")
	}

	*lines = append(*lines, "  ")

	if children, ok := comment["children"].([]interface{}); ok {
		for _, child := range children {
			appendComment(child.(map[string]interface{}), lines, level+1)
		}
	}
}

func wrapText(text, indent string) []string {
	words := strings.Fields(text)
	var lines []string
	var sb strings.Builder

	maxWidth := 80

	for _, word := range words {
		if sb.Len()+len(word)+1 > maxWidth {
			lines = append(lines, indent+sb.String())
			sb.Reset()
		}
		if sb.Len() > 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(word)
	}
	if sb.Len() > 0 {
		lines = append(lines, indent+sb.String())
	}

	return lines
}
