package main

import (
	"reflect"

	"github.com/google/uuid"
)

type EventName string

type Tube struct {
	payloads   map[EventName]reflect.Type
	registarts map[string]Registar
}

func (t *Tube) NewRegistar() (Registar, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	reg := &registar{
		Id: id.String(),
	}
	t.register(reg)
	return reg, nil
}

func (t *Tube) register(r *registar) {
	t.registarts[r.Id] = r
}

// type SyncTube struct{}

// type AsyncTube struct{}
