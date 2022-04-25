package main

import (
	"fmt"
	"time"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern


*/
// CPU Структура ЦП
type CPU struct {
	temperature int
	frequency   float32
}

func (c *CPU) Execute() {
	fmt.Printf("CPU starting execute tasks.\n")
	time.Sleep(time.Second)

}

// Overclock Разгон ЦП
func (c *CPU) Overclock() {
	fmt.Printf("Overclocking cpu.\n")
	c.temperature += 20
	c.frequency += 700
	time.Sleep(time.Second)
	fmt.Printf("CPU frequency has been increased\n")
}

// Fan Кулер ЦП
type Fan struct {
	frequency int
	*CPU
}

// CheckTemperature Проверка температуры ЦП
func (c *Fan) CheckTemperature() {
	if c.temperature >= 90 {
		fmt.Printf("CPU is overheating.Increasing CPU fan frequency\n")
		time.Sleep(time.Second)
		c.IncreaseSpeed()
	} else {
		fmt.Printf("Normal temperature.No actions needed\n")

	}
}

// IncreaseSpeed Увелечение скорости вентиляторов
func (c *Fan) IncreaseSpeed() {
	c.frequency = 2000
	time.Sleep(time.Second)
	fmt.Printf("CPU fan frequency has been increased\n")

}

type RAM struct {
}

func (r *RAM) Load() {
	time.Sleep(time.Second)
	fmt.Printf("Ram loading data\n")
}

type SSD struct{}

func (hd *SSD) Read() {
	time.Sleep(time.Second)
	fmt.Printf("Reading data from ssd\n")
}

// ComputerFacade Структура фасада
type ComputerFacade struct {
	cpu *CPU
	fan *Fan
	ram *RAM
	ssd *SSD
}

// NewComputerFacade Констркуктор
func NewComputerFacade() *ComputerFacade {
	f := new(ComputerFacade)
	f.cpu = &CPU{
		temperature: 70,
		frequency:   3500,
	}
	f.fan = &Fan{
		frequency: 1000,
		CPU:       f.cpu,
	}
	f.ram = new(RAM)
	f.ssd = new(SSD)
	return f
}

// Start Фасад,скрывающий внутренние операции
func (c *ComputerFacade) Start() {
	c.cpu.Execute()
	c.fan.CheckTemperature()
	c.ram.Load()
	c.ssd.Read()
	time.Sleep(time.Second)
	fmt.Printf("Computer successfully started\n")

}
func main() {
	computer := NewComputerFacade()
	computer.Start()
	computer.cpu.Overclock()
	computer.fan.CheckTemperature()
}
