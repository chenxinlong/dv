package dv

import (
	"os"
	"strings"

	"github.com/chenxinlong/dv/configs"
	"github.com/chenxinlong/dv/internal/app/dv/handler"

	aw "github.com/deanishe/awgo"
)

var (
	wf       *aw.Workflow
	handlers handler.HandlerContainer
)

func init() {
	if err := configs.LoadConfig(); err != nil {
		panic(err)
	}

	wf = aw.New()

	handlers = handler.NewHandlerContainer(wf)
	handlers.Register(handler.HandlerGithub, handler.NewHandlerGithub())
	handlers.Register(handler.HandlerGoogle, handler.NewHandlerGoogle())
	handlers.Register(handler.HandlerWikipedia, handler.NewHandlerWikipedia())
	handlers.Register(handler.HandlerStackOverFlow, handler.NewHandlerStackOverFlow())
	handlers.Register(handler.HandlerDeepl, handler.NewHandlerDeepl())
	handlers.Register(handler.HandlerIP, handler.NewHandlerIP())
	handlers.Register(handler.HandlerWeather, handler.NewHandlerWeather())
}

func Run() {
	wf.Run(proc)
}

func proc() {
	defer func() {
		wf.SendFeedback()
	}()

	arg := os.Args[1]
	if strings.TrimSpace(arg) == "" {
		wf.NewItem("type something...")
		return
	}

	handlers.Handle(arg)
}
