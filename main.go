package main

import (
	"fmt"
	"time"
)

func main() {

	githubToken := GetValueFromEnv("GITHUB_TOKEN")
	githubToken = "0471fc4285b2db726d8fdb85adf0dc80d2c5379c"

	t := time.Tick(time.Hour * 1)
	isFirst := true
	var today string

	for {
		/* 爬虫获取新闻 */
		var content string
		current := time.Now().Format("2006-01-02")

		if isFirst || current != today {

			pushTime := time.Now()
			err, contentList := GetNewsContent(pushTime, "1")

			if err != nil {
				fmt.Printf("get newsList err:%v", err)
			} else {
				isFirst = false
				today = time.Now().Format("2006-01-02")
				for _, c := range contentList {
					content = content + c
				}
			}
			/* 推送消息 */
			if content != "" {
				githubContent := ""
				for _, c := range contentList {
					githubContent = githubContent + "- " + c
				}

				err := PushGithub(githubToken, pushTime, githubContent)

				if err != nil {
					fmt.Printf("push to github err:%v", err.Error())
				}

			}

		}
		<-t
	}

}
