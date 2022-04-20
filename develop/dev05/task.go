package main

import (
	"bufio"
	"flag"
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

func main() {
	var (
		A int
		B int
		C int
		c bool
		i bool
		v bool
		F bool
		n bool
	)
	lines := make([]string, 0)

	flag.IntVar(&A, "A", 0, "Печать +N строк после совпадения")
	flag.IntVar(&B, "B", 0, "Печать +N строк до совпадения")
	flag.IntVar(&C, "C", 0, "Печать +N строк вокруг совпадения")
	flag.BoolVar(&c, "c", false, "Печать количества совпадений")
	flag.BoolVar(&i, "i", false, "Игнорирование регистра")
	flag.BoolVar(&v, "v", false, "Исключать строки с совпадениями")
	flag.BoolVar(&F, "F", false, "Точное совпадение со строкой")
	flag.BoolVar(&n, "n", false, "Печать номер строки")
	flag.Parse()

	if C > B {
		B = C
	}
	if C > A {
		A = C
	}

	pattern := flag.Args()[0]

	input, _ := os.Open("test.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

}
func getIndexes(lines []string, pattern string) []int {
	indexes := make([]int, 0)
	for i, st := range lines {
		if strings.Contains(st, pattern) {
			indexes = append(indexes, i)
		}
	}
	return indexes
}
func after(lines []string, pattern string, N int) []string {
	res := make([]string, 0)
	indexes := getIndexes(lines, pattern)
	for _, i := range indexes {
		for j := i; j < i+N; j++ {
			res = append(res, lines[i])
		}
	}
	return res
}
func before(lines []string, pattern string, N int) []string {
	res := make([]string, 0)
	indexes := getIndexes(lines, pattern)
	for _, i := range indexes {
		for j := i; j < i+N; j++ {
			res = append(res, lines[i])
		}
	}
	return res
}
func countMatches(lines []string, pattern string) int {
	return len(getIndexes(lines, pattern))
}
func ignore(lines []string, pattern string) []string {
	lowPattern := strings.ToLower(pattern)
	res := make([]string, 0)
	for _, st := range lines {
		lowSt := strings.ToLower(st)
		if strings.Contains(lowSt, lowPattern) {
			res = append(res, st)
		}
	}
	return res
}
