package main

import (
	"fmt"
)

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

type PCparts interface {
	Accept(PCVisitor)
}

type CPUpart struct {
	Name string
}

func (c *CPUpart) Accept(visitor PCVisitor) {
	visitor.visitCPU(c)
}

type GPUpart struct {
	Name string
}

func (g *GPUpart) Accept(visitor PCVisitor) {
	visitor.visitGPU(g)
}

type PersonalComputer struct {
	parts []PCparts
}

func NewPC() *PersonalComputer {
	pc := new(PersonalComputer)
	pc.parts = []PCparts{
		&CPUpart{"Intel"},
		&GPUpart{"Nvidia"},
	}
	return pc
}

func (pc *PersonalComputer) Accept(visitor PCVisitor) {
	for _, part := range pc.parts {
		part.Accept(visitor)
	}
}

// PCVisitor - интрефейс посетителя
type PCVisitor interface {
	visitCPU(cpu *CPUpart)
	visitGPU(engine *GPUpart)
}

// реализация посетителя

type GetMessageVisitor struct {
	Messages []string
}

func (g *GetMessageVisitor) visitCPU(cpu *CPUpart) {
	g.Messages = append(g.Messages, fmt.Sprintf("Visiting %s cpu\n", cpu.Name))
}

func (g *GetMessageVisitor) visitGPU(gpu *GPUpart) {
	g.Messages = append(g.Messages, fmt.Sprintf("Visiting %s gpu\n", gpu.Name))
}

func main() {
	pc := NewPC()
	visitor := new(GetMessageVisitor)
	pc.Accept(visitor)
	fmt.Println(visitor.Messages)
}
