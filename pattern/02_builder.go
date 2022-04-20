package main

import "fmt"

type PcBuilder interface {
	installCPU(string) PcBuilder
	installGPU(string) PcBuilder
	installRAM(string) PcBuilder
	installPSU(string) PcBuilder
	installMotherboard(string) PcBuilder
	installFan(string) PcBuilder
	NewSimplePC() PC
	NewGamingPC() PC
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
	pc.installGPU("Simple GPU")
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
func NewPC() PcBuilder {
	return &PCBuilder{}
}
func main() {
	builder := NewPC()
	a := builder.NewSimplePC()
	b := builder.NewGamingPC()
	fmt.Println(a, b)
}
