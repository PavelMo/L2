package main

import (
	"fmt"
	"time"
)

//Структура ЦП
type CPU struct {
	temperature int
	frequency   float32
}

func (c *CPU) Execute() {
	fmt.Println("CPU starting execute tasks.\n")
	time.Sleep(time.Second)

}

//Разгон ЦП
func (c *CPU) Overclock() {
	fmt.Println("Overclocking cpu.\n")
	c.temperature += 20
	c.frequency += 700
	time.Sleep(time.Second)
	fmt.Println("CPU frequency has been increased\n")
}

//Кулер ЦП
type CPUfan struct {
	frequency int
	*CPU
}

//Проверка температуры ЦП
func (c *CPUfan) CheckTemperature() {
	if c.temperature >= 90 {
		fmt.Println("CPU is overheating.Increasing CPU fan frequency\n")
		time.Sleep(time.Second)
		c.IncreaseSpeed()
	} else {
		fmt.Println("Normal temperature.No actions needed\n")

	}
}

//Увелечение скорости вентиляторов
func (c *CPUfan) IncreaseSpeed() {
	c.frequency = 2000
	time.Sleep(time.Second)
	fmt.Println("CPU fan frequency has been increased\n")

}

type RAM struct {
}

func (r *RAM) Load() {
	time.Sleep(time.Second)
	fmt.Println("Ram loading data\n")
}

type SSD struct{}

func (hd *SSD) Read() {
	time.Sleep(time.Second)
	fmt.Println("Reading data from ssd\n")
}

//Структура фасада
type ComputerFacade struct {
	cpu *CPU
	fan *CPUfan
	ram *RAM
	ssd *SSD
}

//Констркуктор
func NewComputerFacade() *ComputerFacade {
	f := new(ComputerFacade)
	f.cpu = &CPU{
		temperature: 70,
		frequency:   3500,
	}
	f.fan = &CPUfan{
		frequency: 1000,
		CPU:       f.cpu,
	}
	f.ram = new(RAM)
	f.ssd = new(SSD)
	return f
}

//Фасад,скрывающий внутренние операции
func (c *ComputerFacade) Start() {
	c.cpu.Execute()
	c.fan.CheckTemperature()
	c.ram.Load()
	c.ssd.Read()
	time.Sleep(time.Second)
	fmt.Println("Computer successfully started\n")

}
func main() {
	computer := NewComputerFacade()
	computer.Start()
	computer.cpu.Overclock()
	computer.fan.CheckTemperature()
}
