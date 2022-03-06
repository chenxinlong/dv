// Package observer provide a observer pattern implementation
// wip :
// 1. TriggerAndWait()
// 2. add test
package observer

import (
	"reflect"
	"sync"
)

type Callback func()

type Observer interface {
	Watch(event string, cb Callback)
	Unwatch(event string, cb Callback)
	Trigger(event string)
}

type observer struct {
	store sync.Map
}

func (o *observer) Watch(event string, cb Callback) {
	var final []Callback
	if cur, ok := o.store.Load(event); ok {
		final = append(final, cur.([]Callback)...)
	}

	final = append(final, cb)
	o.store.Store(event, final)

	return
}

func (o *observer) Unwatch(event string, cb Callback) {
	var final []Callback

	// compare func using pointer
	pcb := reflect.ValueOf(cb).Pointer()
	if cur, ok := o.store.Load(event); ok {
		for _, value := range cur.([]Callback) {
			if pcb != reflect.ValueOf(value).Pointer() {
				final = append(final, value)
			}
		}
	}

	final = append(final, cb)
	o.store.Store(event, final)

	return
}

// Trigger triggers the callbacks of event one by one
// wip :
// 1. error handle
// 2. TriggerAndWait() --- type Callback(event *Event)
func (o *observer) Trigger(event string) {
	cur, ok := o.store.Load(event)
	if !ok {
		return
	}

	for _, cb := range cur.([]Callback) {
		// TODO : cb() error 处理
		cb()
	}
}
