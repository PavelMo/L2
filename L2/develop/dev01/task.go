package main

import (
	"develop/current_time"
	"fmt"
	"log"
	"os"
)

/*
=== Базовая задача ===
Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.
Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

func main() {
	currTime, err := current_time.GetCurrTime()
	if err != nil {
		//Если возникают ошибки, выводим их в STDERR
		log.Println("Error occurred while getting current time:", err)
		os.Exit(1)
	}
	fmt.Println("Get current time:", currTime)
}
