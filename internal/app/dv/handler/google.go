package handler

import (
	"github.com/chenxinlong/dv/internal/pkg/project"

	aw "github.com/deanishe/awgo"
)

const (
	HandlerGoogle = "google"
)

type GoogleHandler struct {
	icon *aw.Icon
	typ  HandlerType
}

func NewHandlerGoogle() *GoogleHandler {
	return &GoogleHandler{
		icon: &aw.Icon{
			Value: project.PROJDIR + "/static/icons/google.png",
		},
		typ: HandlerTypFacade,
	}
}

func (g GoogleHandler) GetType() HandlerType {
	return g.typ
}

func (g GoogleHandler) Handle(e Event) (item *aw.Item) {
	if e.lv == 0 {
		return
	}
	if e.lv == 1 {
		url := "https://google.com/search?q=" + e.input
		item = e.wf.NewItem("google : " + e.input)
		item.Subtitle(url)
		item.Icon(g.icon)
		item.Valid(true)
		item.Arg(url)
	}

	return
}
