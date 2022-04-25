package main

import (
	"fmt"
)

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

// Context Контекст который меняет стратегию
type Context struct {
	strategy func()
}

func (c *Context) ExecuteStrategy() {
	c.strategy()
}

func (c *Context) SetStrategy(strategy func()) {
	c.strategy = strategy
}

func main() {
	FirstAlgorithm := func() {
		fmt.Println("Running first algorithm")
	}
	SecondAlgorithm := func() {
		fmt.Println("Running second algorithm")
	}
	context := Context{FirstAlgorithm}
	context.ExecuteStrategy()
	context.SetStrategy(SecondAlgorithm)
	context.ExecuteStrategy()
}
