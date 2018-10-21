package main

import "github.com/google/uuid"

// Registar provides the implementer with capability to register callbacks for events
// type Registar interface {
// 	Listen(ev EventName, f callback)
// }

type registar struct {
	ID     string
	events map[EventName]callback
	tube   Tuber
}

type callback func(payload interface{})

// NewRegistar will create a new registar for a caller to use
// An id will be asigned to the current caller
func NewRegistar(t Tuber) (*registar, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	reg := &registar{
		ID:     id.String(),
		events: make(map[EventName]callback),
	}
	reg.tube = t
	reg.tube.Register(reg)
	return reg, nil
}

// IsListeningTo will say if the current registar is listening to the passed event
func (r *registar) IsListeningTo(ev EventName) bool {
	if _, ok := r.events[ev]; ok {
		return ok
	}
	return false
}

// Listen will tell the registar that when such an event occurs and the registar is notified
// to call the provided callback
func (r *registar) Listen(ev EventName, f callback) {
	r.events[ev] = f
}

func (r *registar) SayForAll(ev EventName, payload interface{}) {
	r.tube.Propagate("", ev, payload)
}

func (r *registar) SayForOne(regID string, ev EventName, payload interface{}) {
	// TODO will send for particular registar
}
