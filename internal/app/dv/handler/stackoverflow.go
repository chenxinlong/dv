package handler

import (
	"github.com/chenxinlong/dv/internal/pkg/project"

	aw "github.com/deanishe/awgo"
)

const (
	HandlerStackOverFlow = "stackoverflow"
)

type StackOverFlowHandler struct {
	icon *aw.Icon
	typ  HandlerType
}

func NewHandlerStackOverFlow() *StackOverFlowHandler {
	return &StackOverFlowHandler{
		icon: &aw.Icon{
			Value: project.PROJDIR + "/static/icons/stackoverflow.png",
		},
		typ: HandlerTypFacade,
	}
}

func (g StackOverFlowHandler) GetType() HandlerType {
	return g.typ
}

func (g StackOverFlowHandler) Handle(e Event) (item *aw.Item) {
	if e.lv == 0 {
		return
	}
	if e.lv == 1 {
		url := "https://stackoverflow.com/search?q=" + e.input
		item = e.wf.NewItem("stackoverflow : " + e.input)
		item.Subtitle(url)
		item.Icon(g.icon)
		item.Valid(true)
		item.Arg(url)
	}

	return
}
