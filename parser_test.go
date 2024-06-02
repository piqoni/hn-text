package main

import (
	"testing"
)

func TestParseArticles(t *testing.T) {
	htmlContent := `
<table>
    <tr class="athing">
        <td align="right" valign="top" class="title"><span class="rank">1.</span></td>
        <td valign="top" class="votelinks">
            <center><a id='up_40477653' href='vote?id=40477653&amp;how=up&amp;goto=news'>
                    <div class='votearrow' title='upvote'></div>
                </a></center>
        </td>
        <td class="title"><span class="titleline"><a
                    href="https://newscenter.lbl.gov/2012/05/16/majorana-demonstrator/">Majorana, the search for the most elusive neutrino of all</a><span class="sitebit comhead"> (<a href="from?site=lbl.gov"><span
                            class="sitestr">lbl.gov</span></a>)</span></span></td>
    </tr>
    <tr>
        <td colspan="2"></td>
        <td class="subtext"><span class="subline">
                <span class="score" id="score_40477653">23 points</span> by <a href="user?id=bilsbie"
                    class="hnuser">bilsbie</a> <span class="age" title="2024-05-25T20:35:59"><a
                        href="item?id=40477653">2 hours ago</a></span> <span id="unv_40477653"></span> | <a
                    href="hide?id=40477653&amp;goto=news">hide</a> | <a href="item?id=40477653">1&nbsp;comment</a>
            </span>
        </td>
    </tr>
    <tr class="athing">
        <td align="right" valign="top" class="title"><span class="rank">2.</span></td>
        <td valign="top" class="votelinks">
            <center><a id='up_40474712' href='vote?id=40474712&amp;how=up&amp;goto=news'>
                    <div class='votearrow' title='upvote'></div>
                </a></center>
        </td>
        <td class="title"><span class="titleline"><a
                    href="https://reverse.put.as/2024/05/24/abusing-go-infrastructure/">Abusing Go&#x27;s Infrastructure</a><span class="sitebit comhead"> (<a href="from?site=put.as"><span
                            class="sitestr">put.as</span></a>)</span></span></td>
    </tr>
    <tr>
        <td colspan="2"></td>
        <td class="subtext"><span class="subline">
                <span class="score" id="score_40474712">298 points</span> by <a href="user?id=efge"
                    class="hnuser">efge</a> <span class="age" title="2024-05-25T12:50:00"><a href="item?id=40474712">10
                        hours ago</a></span> <span id="unv_40474712"></span> | <a
                    href="hide?id=40474712&amp;goto=news">hide</a> | <a href="item?id=40474712">62&nbsp;comments</a>
            </span>
        </td>
    </tr>
</table>`

	expectedArticles := []Article{
		{Title: "Majorana, the search for the most elusive neutrino of all", Link: "https://newscenter.lbl.gov/2012/05/16/majorana-demonstrator/", Comments: 1, CommentsLink: "item?id=40477653"},
		{Title: "Abusing Go's Infrastructure", Link: "https://reverse.put.as/2024/05/24/abusing-go-infrastructure/", Comments: 62, CommentsLink: "item?id=40474712"},
	}

	articles, err := parseArticles(htmlContent)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(articles) != len(expectedArticles) {
		t.Fatalf("Expected %d articles, got %d", len(expectedArticles), len(articles))
	}

	for i, article := range articles {
		if article.Title != expectedArticles[i].Title {
			t.Errorf("Expected title %q, got %q", expectedArticles[i].Title, article.Title)
		}
		if article.Link != expectedArticles[i].Link {
			t.Errorf("Expected link %q, got %q", expectedArticles[i].Link, article.Link)
		}
		if article.Comments != expectedArticles[i].Comments {
			t.Errorf("Expected comments %d, got %d", expectedArticles[i].Comments, article.Comments)
		}
		if article.CommentsLink != expectedArticles[i].CommentsLink {
			t.Errorf("Expected comments link %q, got %q", expectedArticles[i].CommentsLink, article.CommentsLink)
		}
	}
}

// func TestParseArticles(t *testing.T) {
// 	t.Run("Empty HTML", func(t *testing.T) {
// 		articles, err := parseArticles("")
// 		assert.NoError(t, err)
// 		assert.Empty(t, articles)
// 	})

// 	t.Run("Valid HTML", func(t *testing.T) {
// 		html := `
// 			<tr class="athing">
// 				<td class="title">
// 					<span class="titleline">
// 						<a href="https://example.com">Article 1</a>
// 					</span>
// 				</td>
// 			</tr>
// 			<tr class="athing">
// 				<td class="title">
// 					<span class="titleline">
// 						<a href="https://example.com/2">Article 2</a>
// 					</span>
// 				</td>
// 			</tr>
// 		`
// 		articles, err := parseArticles(html)
// 		assert.NoError(t, err)
// 		assert.Len(t, articles, 2)
// 		assert.Equal(t, "Article 1", articles[0].Title)
// 		assert.Equal(t, "https://example.com", articles[0].Link)
// 		assert.Equal(t, "Article 2", articles[1].Title)
// 		assert.Equal(t, "https://example.com/2", articles[1].Link)
// 	})

// 	t.Run("Error in goquery", func(t *testing.T) {
// 		doc := &goquery.Document{}
// 		// doc.SetError(fmt.Errorf("test error"))
// 		articles, err := parseArticlesFromDocument(doc)
// 		assert.Error(t, err)
// 		assert.Nil(t, articles)
// 	})
// }

// func parseArticlesFromDocument(doc *goquery.Document) ([]Article, error) {
// 	var articles []Article

// 	doc.Find("tr.athing").Each(func(i int, s *goquery.Selection) {
// 		title := s.Find("td.title > span.titleline > a").Text()
// 		link, _ := s.Find("td.title > span.titleline > a").Attr("href")
// 		commentsCount, err := extractCommentsCount(s)
// 		if err != nil {
// 			commentsCount = 0
// 		}

// 		article := Article{
// 			Title:    title,
// 			Link:     link,
// 			Comments: commentsCount,
// 		}

// 		articles = append(articles, article)
// 	})

// 	return articles, nil
// }
