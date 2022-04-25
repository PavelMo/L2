package main

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/
//director - структура директора, который определяет, что строить
type director struct {
	PCBuilder
	builderType string
}

func (d *director) getComputer(computerType string) PC {
	switch computerType {
	case "simple":
		return d.NewSimplePC()
	case "gaming":
		return d.NewGamingPC()
	}
	return PC{}
}

// PcBuilder - интерфейс для всех строителей
type PcBuilder interface {
	installCPU(string) PcBuilder
	installGPU(string) PcBuilder
	installRAM(string) PcBuilder
	installPSU(string) PcBuilder
	installMotherboard(string) PcBuilder
	installFan(string) PcBuilder
}
type PCBuilder struct {
	CPU         string
	GPU         string
	RAM         string
	PSU         string
	MotherBoard string
	Fan         string
}
type PC struct {
	CPU         string
	GPU         string
	RAM         string
	PSU         string
	MotherBoard string
	Fan         string
}

func (pc *PCBuilder) installCPU(spec string) PcBuilder {
	pc.CPU = spec
	return pc
}
func (pc *PCBuilder) installGPU(spec string) PcBuilder {
	pc.GPU = spec
	return pc
}
func (pc *PCBuilder) installRAM(spec string) PcBuilder {
	pc.RAM = spec
	return pc
}
func (pc *PCBuilder) installPSU(spec string) PcBuilder {
	pc.PSU = spec
	return pc
}
func (pc *PCBuilder) installMotherboard(spec string) PcBuilder {
	pc.MotherBoard = spec
	return pc
}
func (pc *PCBuilder) installFan(spec string) PcBuilder {
	pc.Fan = spec
	return pc
}
func (pc *PCBuilder) NewSimplePC() PC {
	pc.installCPU("Simple CPU")
	pc.installRAM("Simple RAM")
	pc.installPSU("Simple PSU")
	pc.installMotherboard("Simple Motherboard")
	pc.installFan("Simple Fan")
	return *pc.PCspec()
}
func (pc *PCBuilder) NewGamingPC() PC {
	pc.installCPU("Gaming CPU")
	pc.installGPU("Gaming GPU")
	pc.installRAM("Gaming RAM")
	pc.installPSU("Gaming PSU")
	pc.installMotherboard("Gaming Motherboard")
	pc.installFan("Gaming Fan")
	return *pc.PCspec()

}
func (pc *PCBuilder) PCspec() *PC {
	return &PC{
		CPU:         pc.CPU,
		GPU:         pc.GPU,
		RAM:         pc.RAM,
		PSU:         pc.PSU,
		MotherBoard: pc.MotherBoard,
		Fan:         pc.Fan,
	}
}
func main() {
	dir := new(director)
	a := dir.getComputer("simple")
	b := dir.getComputer("gaming")
	fmt.Println(a, b)
}
