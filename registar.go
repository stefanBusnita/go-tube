package main

type Registar interface {
	Listen(ev EventName, f callback)
}

type registar struct {
	Id string
}

type callback func(payload interface{})

func (r *registar) Listen(ev EventName, f callback) {}
