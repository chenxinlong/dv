package handler

import (
	"strconv"

	"github.com/chenxinlong/dv/internal/pkg/project"

	aw "github.com/deanishe/awgo"
)

const (
	HandlerDeepl = "deepl"
)

type DeeplHandler struct {
	icon *aw.Icon
	typ  HandlerType
}

func NewHandlerDeepl() *DeeplHandler {
	return &DeeplHandler{
		icon: &aw.Icon{
			Value: project.PROJDIR + "/static/icons/deepl.png",
		},
		typ: HandlerTypFacade,
	}
}

func (g DeeplHandler) GetType() HandlerType {
	return g.typ
}

func (g DeeplHandler) Handle(e Event) (item *aw.Item) {
	if e.lv == 0 {
		return
	}
	if e.lv == 1 {
		item = e.wf.NewItem("deepl : " + e.input)
		item.Subtitle("https://deepl.com/search?q=" + e.input)
		item.Icon(g.icon)
		item.Valid(true)
		item.Arg("dv deepl:" + e.input)
	}
	if e.lv == 2 {
		// debug
		for i := 0; i < 3; i++ {
			item = e.wf.NewItem(strconv.Itoa(i) + " input=" + e.input)
			item.Icon(g.icon)
		}
		return
	}

	return
}
