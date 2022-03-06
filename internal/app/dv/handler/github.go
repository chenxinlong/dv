package handler

import (
	"context"
	"fmt"

	"github.com/chenxinlong/dv/internal/pkg/project"

	aw "github.com/deanishe/awgo"
	"github.com/google/go-github/v42/github"
)

const (
	HandlerGithub = "github"
)

type GithubHandler struct {
	icon *aw.Icon
	typ  HandlerType
}

func NewHandlerGithub() *GithubHandler {
	return &GithubHandler{
		icon: &aw.Icon{
			Value: project.PROJDIR + "/static/icons/github.png",
		},
		typ: HandlerTypFacade,
	}
}

func (g GithubHandler) GetType() HandlerType {
	return g.typ
}

func (g GithubHandler) Handle(e Event) (item *aw.Item) {
	if e.lv == 0 {
		return
	}
	if e.lv == 1 {
		item = e.wf.NewItem("github : " + e.input)
		item.Subtitle("https://github.com/search?q=" + e.input)
		item.Icon(g.icon)
		item.Valid(true)
		item.Arg("dv github:" + e.input)
	}
	if e.lv == 2 {
		e.wf.NewItem(fmt.Sprintf("Search %q on github.com", e.rawInput)).Valid(true).Arg("https://github.com/search?q=" + e.rawInput)

		// search from github
		cli := github.NewClient(nil)
		result, _, err := cli.Search.Repositories(context.Background(), e.rawInput, &github.SearchOptions{
			ListOptions: github.ListOptions{PerPage: 10},
		})
		if err != nil {
			e.wf.NewItem("failed to search from github, err = " + err.Error())
			return
		}

		// construct items
		for _, repo := range result.Repositories {
			lang := repo.GetLanguage()
			if lang == "" {
				lang = "unknown"
			}

			title := fmt.Sprintf("️☆:%d [%s] %s", repo.GetStargazersCount(), lang, repo.GetFullName())
			e.wf.NewItem(title).
				Icon(g.icon).
				Subtitle(repo.GetHTMLURL()).
				Valid(true).
				Arg(repo.GetHTMLURL())
		}
		return
	}

	return
}
