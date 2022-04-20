package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
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
	currTime, err := getCurrTime()
	if err != nil {
		//Если возникают ошибки, выводим их в STDERR
		_, err = fmt.Fprintln(os.Stderr, "Error occurred while getting current time:", err)
		if err != nil {
			fmt.Println(err)
		}
		os.Exit(1)
	}
	fmt.Println("Get current time:", currTime)
}

//Получаем время с NTP сервера и в случае ошибки возвращаем её
func getCurrTime() (time.Time, error) {
	currTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		return time.Time{}, err
	}
	return currTime, nil
}
