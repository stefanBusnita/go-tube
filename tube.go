package main

import (
	"reflect"
)

type EventName string

type Tuber interface {
	Register(r *registar)
	Propagate(regID string, ev EventName, payload interface{})
}

// NewTuber will create a new Tube ready for use
func NewTuber() Tuber {
	return &tube{
		payloads:  make(map[EventName]reflect.Type),
		registars: make(map[string]*registar),
	}
}

type tube struct {
	payloads  map[EventName]reflect.Type
	registars map[string]*registar
}

// Register will let the tube know that the registar wants to be part of the channel
func (t *tube) Register(r *registar) {
	t.registars[r.ID] = r
}

// Propagate will propagate the given event with the payload to all registars that listen to it
func (t *tube) Propagate(regID string, ev EventName, payload interface{}) {
	for _, registar := range t.registars {
		if registar.IsListeningTo(ev) {
			registar.events[ev](payload)
		}
	}
}
