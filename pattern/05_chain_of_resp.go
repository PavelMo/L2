package main

import (
	"fmt"
)

type Handler interface {
	Request(flag bool)
}

type HandlerA struct {
	next Handler
}

func (h *HandlerA) Request(flag bool) {
	fmt.Println("Get request in HandlerA")
	if flag {
		h.next.Request(flag)
	}
}

type HandlerB struct {
	next Handler
}

func (h *HandlerB) Request(flag bool) {
	fmt.Println("Get request in HandlerB")
}

func main() {
	handlerA := &HandlerA{new(HandlerB)}
	handlerA.Request(true)
}
