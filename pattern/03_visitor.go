package main

import (
	"fmt"
)

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

func NewCar() *PersonalComputer {
	pc := new(PersonalComputer)
	pc.parts = []PCparts{
		&CPUpart{"Some CPU"},
		&GPUpart{"Some GPU"},
	}
	return pc
}

func (pc *PersonalComputer) Accept(visitor PCVisitor) {
	for _, part := range pc.parts {
		part.Accept(visitor)
	}
}

//Interface of the visitor
type PCVisitor interface {
	visitCPU(cpu *CPUpart)
	visitGPU(engine *GPUpart)
}

//Concrete Implementation of the visitor
type GetMessageVisitor struct {
	Messages []string
}

func (this *GetMessageVisitor) visitCPU(cpu *CPUpart) {
	this.Messages = append(this.Messages, fmt.Sprintf("Visiting the %v wheel\n", cpu.Name))
}

func (this *GetMessageVisitor) visitGPU(gpu *GPUpart) {
	this.Messages = append(this.Messages, fmt.Sprintf("Visiting engine\n"))
}

//Usage of the visitor
func main() {
	car := NewCar()
	visitor := new(GetMessageVisitor)
	car.Accept(visitor)
	fmt.Println(visitor.Messages)
}
