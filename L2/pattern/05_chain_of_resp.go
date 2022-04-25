package main

import (
	"fmt"
)

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

//Запрос с номером нужного обработчика и статусом
type req struct {
	done   bool
	needle int
}

// Handler - Абстрактный обработчик с двумя методами
type Handler interface {
	Request(*req)
	Next(Handler)
}

//Первый обработчик
type firstHandler struct {
	next Handler
}

func (h *firstHandler) Request(r *req) {
	fmt.Println("Got request in first handler")
	if r.done {
		fmt.Println("Request already processed")
	} else if r.needle == 1 {
		r.done = true
		fmt.Printf("Request processed in first handler...\n Done...\n")
		return
	}
	h.next.Request(r)
}
func (h *firstHandler) Next(next Handler) {
	h.next = next
}

//Второй обработчик
type secondHandler struct {
	next Handler
}

func (h *secondHandler) Request(r *req) {
	fmt.Println("Get request in second handler")
	if r.done {
		fmt.Println("Request already processed")
	} else if r.needle == 2 {
		r.done = true
		fmt.Printf("Request processed in a second handler...\n Done...\n")
	}
}
func (h *secondHandler) Next(next Handler) {
	h.next = next
}
func newReq(done bool, needle int) *req {
	return &req{
		done:   done,
		needle: needle,
	}
}
func main() {
	r := newReq(false, 2)

	s := &secondHandler{}

	f := &firstHandler{}
	f.Next(s)

	f.Request(r)
}
