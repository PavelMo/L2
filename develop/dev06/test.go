package main

import (
	"bufio"
	"errors"
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
		log.Fatal(errors.New("you need to specify a field"))
	}
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	tmp := strings.Split(f, ",")

	fields := make([]int, len(tmp))

	for i := range tmp {
		num, err := strconv.Atoi(tmp[i])
		if err != nil || num == 0 {
			log.Fatal("cannot convert string to int", err)
		}
		fields[i] = num
	}
	lines = cutSt(lines, d, fields)
	fmt.Println(lines, s, f, d)
	if s {
		res := onlyDelim(lines, d)
		for _, val := range res {
			fmt.Println(val)
		}
	} else {
		for _, val := range lines {
			fmt.Println(val)
		}
	}

}
func onlyDelim(lines []string, delimiter string) []string {
	res := make([]string, 0)
	for i := range lines {
		if strings.Contains(lines[i], delimiter) {
			res = append(res, lines[i])
		}
	}
	lines = res
	return res
}
func cutSt(lines []string, delimiter string, fields []int) []string {
	res := make([]string, 0)
	for _, st := range lines {
		if delimiter != "" && strings.Contains(st, delimiter) {
			columns := strings.Split(st, delimiter)
			resSt := strings.Builder{}
			for _, val := range fields {
				if len(columns) >= val {
					resSt.WriteString(columns[val-1])
					resSt.WriteString(delimiter)
				}
			}
			res = append(res, strings.TrimSuffix(resSt.String(), delimiter))
		}
	}
	return res
}
