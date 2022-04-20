package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===
Отсортировать строки (man sort)
Основное
Поддержать ключи
-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки
Дополнительное
Поддержать ключи
-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	var (
		k   int
		n   bool
		r   bool
		u   bool
		err error
	)

	flag.IntVar(&k, "k", -1, "указание колонки для сортировки")
	flag.BoolVar(&n, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&r, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&u, "u", false, "не выводить повторяющиеся строки")
	flag.Parse()

	var lines = []string{""}
	input, _ := os.Open("test.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if u {
		lines = deleteDuplicates(lines)
	}
	if !n && k > -1 {
		lines = columnSort(k, lines)
	}
	if n {
		lines, err = NumberSort(k, lines)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
	if r {
		for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
			lines[i], lines[j] = lines[j], lines[i]
		}
	}
	for _, val := range lines {
		fmt.Println(val)
	}
	output, err := os.Create("output.txt")
	defer output.Close()

	w := bufio.NewWriter(output)
	for _, str := range lines {
		w.WriteString(str + "\n")
	}

	w.Flush()
}
func getNumFromColumn(n int, st string) (int, error) {
	if len(st) == 0 {
		return 0, errors.New("empty string")
	}
	arrSt := strings.Fields(st)
	if n < len(arrSt) {
		num, err := strconv.Atoi(arrSt[n])
		return num, err
	} else {
		return 0, errors.New("incorrect column definition")
	}
}
func NumberSort(n int, lines []string) ([]string, error) {
	Nums := make([]int, 0)
	saveMp := make(map[int]string)
	for _, st := range lines {
		if st == "" {
			continue
		}
		num, err := getNumFromColumn(n, st)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		saveMp[num] = st
		Nums = append(Nums, num)
	}

	Result := make([]string, 0)
	//fmt.Println(Nums)
	sort.Ints(Nums)
	for _, val := range Nums {
		Result = append(Result, saveMp[val])
	}
	return Result, nil
}
func getWordFromColumn(n int, st string) string {
	if len(st) == 0 {
		return ""
	}
	arrSt := strings.Fields(st)
	if n < len(arrSt) {
		return arrSt[n]
	} else {
		return ""
	}
}
func columnSort(n int, lines []string) []string {
	res := make([]string, 0)
	column := make([]string, 0)
	saveMp := make(map[string]string)

	for _, st := range lines {
		word := getWordFromColumn(n, st)
		if word != "" {
			column = append(column, word)
			saveMp[word] = st
		}
	}
	sort.Strings(column)
	for _, word := range column {
		res = append(res, saveMp[word])
	}
	return res
}
func deleteDuplicates(fileStrings []string) []string {
	unique := make(map[string]struct{})

	for _, st := range fileStrings {
		unique[st] = struct{}{}
	}
	res := make([]string, 0, len(unique))
	for st := range unique {
		res = append(res, st)
	}
	return res
}
