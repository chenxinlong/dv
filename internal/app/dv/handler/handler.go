package handler

import (
	"fmt"
	"strings"
	"sync"

	"golang.org/x/sync/errgroup"

	aw "github.com/deanishe/awgo"
)

const (
	HandlerTypFacade HandlerType = iota
	HandlerTypShortTag
)

type (
	HandlerContainer interface {
		Register(name string, handler Handler)
		UnRegister(name string)
		Get(name string) Handler
		Handle(input string)
	}

	HandlerType uint8

	Handler interface {
		Handle(event Event) *aw.Item
		GetType() HandlerType
	}

	Event struct {
		wf       *aw.Workflow
		lv       uint8
		input    string
		rawInput string
	}
)

type handlerContainerImpl struct {
	wf    *aw.Workflow
	store sync.Map // store stores handler, key=handler name, value=handler instance
}

func NewHandlerContainer(wf *aw.Workflow) HandlerContainer {
	container := &handlerContainerImpl{
		wf:    wf,
		store: sync.Map{},
	}

	return container
}

func (h *handlerContainerImpl) Register(name string, handler Handler) {
	h.store.Store(name, handler)
}

func (h *handlerContainerImpl) UnRegister(name string) {
	h.store.Delete(name)
}

func (h *handlerContainerImpl) Get(name string) Handler {
	val, _ := h.store.Load(name)
	return val.(Handler)
}

func (h *handlerContainerImpl) Handle(input string) {
	if input[0] == '-' {
		h.handleShortTag(input)
	} else {
		_ = h.handleFacade(input)
	}
}

func (h *handlerContainerImpl) handleShortTag(input string) {
	event := Event{
		wf:       h.wf,
		lv:       0,
		input:    input,
		rawInput: input,
	}
	tag := strings.Split(strings.ToLower(strings.TrimPrefix(input, "-")), " ")[0]
	if tag == "i" {
		h.Get(HandlerIP).Handle(event)
	}
	if tag == "w" {
		h.Get(HandlerWeather).Handle(event)
	}
}

func (h *handlerContainerImpl) handleFacade(input string) (err error) {
	g := errgroup.Group{}

	// itemMap store items for each lv1 handler
	// for further item sort
	itemMap := sync.Map{}
	sortKeys := []string{HandlerDeepl, HandlerGithub, HandlerStackOverFlow, HandlerGoogle, HandlerWikipedia}
	reorder := true
	h.store.Range(func(key, value interface{}) bool {
		handler := value.(Handler)
		if handler.GetType() != HandlerTypFacade {
			return true
		}

		// construct event
		event := Event{
			wf:       h.wf,
			lv:       1,
			input:    input,
			rawInput: input,
		}
		if strings.Contains(input, fmt.Sprintf("dv ")) && !strings.Contains(input, fmt.Sprintf("dv %s:", key)) {
			reorder = false
			event.lv = 0
		}
		if strings.Contains(input, fmt.Sprintf("dv %s:", key)) {
			reorder = false
			event.lv = 2
			event.rawInput = strings.Split(input, fmt.Sprintf("dv %s:", key))[1]
		}

		// handle
		g.Go(func() error {
			item := handler.Handle(event)
			itemMap.Store(key.(string), item)
			return nil
		})

		return true
	})
	if err = g.Wait(); err != nil {
		return
	}

	// clear and rebuild sorted items
	if !reorder {
		return
	}
	h.wf.Feedback.Clear()
	for _, key := range sortKeys {
		item, ok := itemMap.Load(key)
		if ok {
			h.wf.Feedback.Items = append(h.wf.Feedback.Items, item.(*aw.Item))
		}
	}

	return
}
