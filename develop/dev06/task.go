package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===
Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные
Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	var (
		f string
		d string
		s bool
	)

	flag.StringVar(&f, "f", "", "Выбор определенной колонки")
	flag.StringVar(&d, "d", "\t", "Использование другого разделителя")
	flag.BoolVar(&s, "s", false, "Выбор строк только с разделителем")
	flag.Parse()

	if len(f) < 1 {
		log.Fatalln("you need to specify a field")
	}
	if len(flag.Args()) < 1 {
		log.Fatalln("file is not specified")
	}
	file, err := os.Open(flag.Args()[0])
	if err != nil {
		log.Fatalln(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	parse := strings.Split(f, ",")
	numFields := make([]int, 0)
	for _, field := range parse {
		digit, err := strconv.Atoi(field)
		if digit == 0 || err != nil {
			log.Fatalln("Incorrect input")
		}
		numFields = append(numFields, digit)
	}
	Result := cut(lines, numFields, d, s)
	for _, st := range Result {
		fmt.Println(st)
	}

}
func cut(lines []string, numFields []int, delim string, sep bool) []string {
	Result := make([]string, 0)
	for _, st := range lines {
		//Делим строку по разделителю
		fields := strings.Split(st, delim)
		if len(fields)-1 > 0 {
			var resSt string
			//Находим нужную колонку
			for i, field := range numFields {
				if len(fields) >= field {
					if i > 0 {
						resSt = resSt + delim + fields[field-1]
					} else {
						resSt = fields[field-1]
					}

				}
			}
			//Записываем в результат
			Result = append(Result, resSt)
		} else if !sep {
			Result = append(Result, st)
		}
	}
	return Result
}
