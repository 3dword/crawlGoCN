package main

import (
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	"time"
)

// 爬虫CoNews网页
func GetNewsContent(publishTime time.Time, day string) (e error, content []string) {

	var baseUrl string

	data := publishTime.Format("2006-01-02")
	dateOther := publishTime.Format("2006-01-2")

	c := colly.NewCollector()
	// Find and visit all links
	c.OnHTML("div > h4 > a", func(e *colly.HTMLElement) {
		fmt.Println(e)
		if strings.Contains(e.Text, data) {
			baseUrl = e.Attr("href")
			fmt.Printf("Link found: %q -> %s\n", e.Text, baseUrl)
		} else if strings.Contains(e.Text, dateOther) {
			baseUrl = e.Attr("href")
			fmt.Printf("Link found: %q -> %s\n", e.Text, baseUrl)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	e = c.Visit("https://gocn.vip/question/sort_type-new__category-14__day-0__page-" + day)
	if e != nil {
		return e, nil
	}

	if baseUrl == "" {
		return errors.New("news not update"), nil
	}


	b := colly.NewCollector()

	// Find and visit all links
	i := 0
	contentList := make([]string, 15)
	b.OnHTML("div.mod-body > div > ol > li", func(e *colly.HTMLElement) {
		if e.Text != "" {
			contentList[i] = TrimQuotes(fmt.Sprintf("%d. %s\n\n", i+1, e.Text))
			i++
			fmt.Printf("%d:%q\n", i, TrimQuotes(fmt.Sprintf("%d. %s\n\n", i+1, e.Text)))
		}
	})
	b.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	e = b.Visit(baseUrl)
	if e != nil {
		return e, nil
	}
	var flag bool
	for _, c := range contentList {
		if c != "" {
			flag = true
			break
		}
	}
	if !flag {
		c := colly.NewCollector()
		c.OnHTML("div.mod-body > div > p", func(e *colly.HTMLElement) {
			if e.Text != "" {
				contentList[i] = TrimQuotes(fmt.Sprintf("%d. %s\n\n", i+1, e.Text))
				i++
				fmt.Printf("%d:%q\n", i, TrimQuotes(fmt.Sprintf("%d. %s\n\n", i+1, e.Text)))
			}
		})
		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL)
		})
		e = c.Visit(baseUrl)
		if e != nil {
			return e, nil
		}
	}
	return nil, contentList

}
