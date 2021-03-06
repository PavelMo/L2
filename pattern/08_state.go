package main

import (
	"fmt"
)

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

type Cont struct {
	state State
}

func (c *Cont) Request() {
	c.state.Handle()
}

func (c *Cont) SetState(state State) {
	c.state = state
}

type State interface {
	Handle()
}

type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle() {
	fmt.Println("ConcreteStateA.Handle()")
}

type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle() {
	fmt.Println("ConcreteStateB.Handle()")
}

func main() {
	context := Cont{new(ConcreteStateA)}
	context.Request()
	context.SetState(new(ConcreteStateB))
	context.Request()
}
