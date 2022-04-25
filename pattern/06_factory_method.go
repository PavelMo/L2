package main

import (
	"fmt"
)

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

// Engine Интерфейс для создания разных двигателей
type Engine interface {
	construct()
}

// CarEngine структура для создания автомобильных двигателей
type CarEngine struct {
}

func (c CarEngine) construct() {
	fmt.Println("Construct car engine")
}

// AirplaneEngine структура для создания самолётных двигателей
type AirplaneEngine struct {
}

func (a AirplaneEngine) construct() {
	fmt.Println("Construct airplane engine")
}

// ConstructEngine Конструктор, который определяет тип двигателя для создания, удовляющему интерфейсу Engine ,двигателя
func ConstructEngine(engineType string) Engine {
	switch engineType {
	case "car":
		return CarEngine{}
	case "airplane":
		return AirplaneEngine{}
	}
	return nil
}
func main() {
	ConstructEngine("airplane").construct()
}
