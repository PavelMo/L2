package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
=== Утилита grep ===
Реализовать утилиту фильтрации (man grep)
Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Args struct {
	A int
	B int
	C int
	c bool
	i bool
	v bool
	F bool
	n bool
}

func main() {

	lines := make([]string, 0)

	A := flag.Int("A", 0, "Печать +N строк после совпадения")
	B := flag.Int("B", 0, "Печать +N строк до совпадения")
	C := flag.Int("C", 0, "Печать +N строк вокруг совпадения")
	c := flag.Bool("c", false, "Печать количество совпадений")
	i := flag.Bool("i", false, "Игнорирование регистра")
	v := flag.Bool("v", false, "Исключать строки с совпадениями")
	F := flag.Bool("F", false, "Точное совпадение со строкой")
	n := flag.Bool("n", false, "Печать номер строки")
	flag.Parse()

	if len(flag.Args()) < 1 {
		log.Println("pattern  is not specified ")
	}
	if len(flag.Args()) < 2 {
		log.Println("source is not specified")
	}
	if *A < 0 || *B < 0 || *C < 0 {
		log.Println("incorrect flag arguments")
	}
	a := &Args{
		A: *A,
		B: *B,
		C: *C,
		c: *c,
		i: *i,
		v: *v,
		F: *F,
		n: *n,
	}

	pattern := flag.Args()[0]
	file := flag.Args()[1]
	input, _ := os.Open(file)
	defer func(input *os.File) {
		err := input.Close()
		if err != nil {

		}
	}(input)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if a.c {
		var count, countInvert int
		for _, st := range lines {
			if strings.Contains(st, pattern) {
				count++
			} else if a.v {
				countInvert++
			}
		}
		if countInvert == 0 {
			fmt.Println(count)
		} else {
			fmt.Println(countInvert)
		}
	} else {
		Result := a.grep(lines, pattern)
		for _, st := range Result {
			fmt.Println(st)
		}
	}

}
func (a *Args) grep(lines []string, pattern string) []string {
	afterMap := make(map[int]struct{})
	beforeMap := make(map[int]struct{})
	res := make([]string, 0)
	if a.A < a.C {
		a.A = a.C
	}
	if a.B < a.C {
		a.B = a.C
	}
	if a.v {
		afterMap = getIndexesAfter(lines, pattern, a.F, a.i, 0)
		for i, st := range lines {
			_, aOk := afterMap[i]
			if !aOk {
				res = append(res, st)
			}
		}
		return res
	}
	afterMap = getIndexesAfter(lines, pattern, a.F, a.i, a.A)
	if a.B > 0 {
		beforeMap = getIndexesBefore(lines, pattern, a.F, a.i, a.B)
	}
	var skipLines bool
	for i, st := range lines {
		_, aOk := afterMap[i]
		_, bOk := beforeMap[i]
		if aOk || bOk {
			if a.n {
				res = append(res, fmt.Sprintf("%d:%s", i+1, st))
				//skipLines = true
			} else {
				res = append(res, st)
				skipLines = true
			}

		} else if skipLines {
			res = append(res, "--")
			skipLines = false
		}
	}
	if res[len(res)-1] == "--" {
		return res[:len(res)-1]
	}
	return res
}

//Записываем в хэш таблицу индексы строк с нужным паттерном/строкой +N строк
func getIndexesAfter(lines []string, needle string, Fixed, Ignore bool, N int) map[int]struct{} {
	indexes := make(map[int]struct{})
	switch {
	case Ignore && Fixed:
		lowNeedle := strings.ToLower(needle)
		for i, st := range lines {
			lowSt := strings.ToLower(st)
			if lowSt == lowNeedle {
				getAfter(indexes, i, N, len(lines))
			}
		}
	case Fixed:
		for i, st := range lines {
			if st == needle {
				getAfter(indexes, i, N, len(lines))
			}
		}
	case Ignore:
		lowNeedle := strings.ToLower(needle)
		for i, st := range lines {
			lowSt := strings.ToLower(st)
			if strings.Contains(lowSt, lowNeedle) {
				getAfter(indexes, i, N, len(lines))
			}
		}
	default:
		for i, st := range lines {
			if strings.Contains(st, needle) {
				getAfter(indexes, i, N, len(lines))
			}
		}
	}
	return indexes
}
func getAfter(indexes map[int]struct{}, i, N, lenLines int) {
	for j := i; j <= i+N && j < lenLines; j++ {
		indexes[j] = struct{}{}
	}
}

//Записываем в хэш таблицу индексы строк с нужным паттерном/строкой -N строк
func getIndexesBefore(lines []string, needle string, Fixed, Ignore bool, N int) map[int]struct{} {
	indexes := make(map[int]struct{})
	switch {
	case Ignore && Fixed:
		lowNeedle := strings.ToLower(needle)
		for i, st := range lines {
			lowSt := strings.ToLower(st)
			if lowSt == lowNeedle {
				getBefore(indexes, i, N)
			}
		}
	case Fixed:
		for i, st := range lines {
			if st == needle {
				getBefore(indexes, i, N)
			}
		}
	case Ignore:
		lowNeedle := strings.ToLower(needle)
		for i, st := range lines {
			lowSt := strings.ToLower(st)
			if strings.Contains(lowSt, lowNeedle) {
				getBefore(indexes, i, N)
			}
		}
	default:
		for i, st := range lines {
			if strings.Contains(st, needle) {
				getBefore(indexes, i, N)
			}
		}
	}
	return indexes
}
func getBefore(indexes map[int]struct{}, i, N int) {
	if i-N >= 0 {
		for j := i - N; j <= i; j++ {
			indexes[j] = struct{}{}
		}
	} else {
		for j := 0; j <= i; j++ {
			indexes[j] = struct{}{}
		}
	}
}
