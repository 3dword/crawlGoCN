package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"log"
	"strings"
	"time"
)

func PushGithub(token string, publish time.Time, contentList string) error {

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	c := "add gocn news--" + publish.Format("2006-01-02")
	sha := ""
	content := &github.RepositoryContentFileOptions{
		Message: &c,
		SHA:     &sha,
		Committer: &github.CommitAuthor{
			Name:  github.String("lubanproj"),
			Email: github.String("1811704358@qq.com"),
			Login: github.String("lubanproj"),
		},
		Author: &github.CommitAuthor{
			Name:  github.String("lubanproj"),
			Email: github.String("1811704358@qq.com"),
			Login: github.String("lubanproj"),
		},
		Branch: github.String("master"),
	}
	op := &github.RepositoryContentGetOptions{}

	repo, _, _, er := client.Repositories.GetContents(ctx, "lubanproj", "go_read", "README.md", op)
	if er != nil || repo == nil {
		fmt.Println(er)
		return er
	}

	content.SHA = repo.SHA
	decodeBytes, err := base64.StdEncoding.DecodeString(*repo.Content)
	if err != nil {
		log.Fatalln(err)
		return err
	}


	oldContentList := strings.Split(string(decodeBytes), "## gocn_news__2019")

	content.Content = []byte(oldContentList[0] + "\n" + "## gocn_news__2019" + "\n" + "### gocn_news_" + publish.Format("2006-01-02") + "\n" + contentList + "\n" + oldContentList[1])

	_, _, err = client.Repositories.UpdateFile(ctx, "lubanproj", "go_read", "README.md", content)
	if err != nil {
		println(err)
		return err
	}
	return nil

}

