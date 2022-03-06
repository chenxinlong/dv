package handler

import (
	"github.com/chenxinlong/dv/internal/pkg/project"

	aw "github.com/deanishe/awgo"
)

const (
	HandlerWikipedia = "wikipedia"
)

type WikipediaHandler struct {
	icon *aw.Icon
	typ  HandlerType
}

func NewHandlerWikipedia() *WikipediaHandler {
	return &WikipediaHandler{
		icon: &aw.Icon{
			Value: project.PROJDIR + "/static/icons/wikipedia.png",
		},
		typ: HandlerTypFacade,
	}
}

func (g WikipediaHandler) GetType() HandlerType {
	return g.typ
}

func (g WikipediaHandler) Handle(e Event) (item *aw.Item) {
	if e.lv == 0 {
		return
	}
	if e.lv == 1 {
		url := "https://en.wikipedia.org/wiki/" + e.input
		item = e.wf.NewItem("wikipedia : " + e.input)
		item.Subtitle(url)
		item.Icon(g.icon)
		item.Valid(true)
		item.Arg(url)
	}

	return
}
