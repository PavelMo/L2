package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
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
	flag.IntVar(&k, "k", 1, "указание колонки для сортировки")
	flag.BoolVar(&n, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&r, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&u, "u", false, "не выводить повторяющиеся строки")
	flag.Parse()

	var lines = make([]string, 0)
	if len(flag.Args()) < 1 {
		log.Println("no files to sort")
		os.Exit(1)
	}
	for _, val := range flag.Args() {
		input, _ := os.Open(val)
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		closeFile(input)
	}
	if u {
		lines = deleteDuplicates(lines, k-1)
	}
	if !n {
		lines = columnSort(k-1, lines)
	}
	if n {
		lines, err = NumberSort(k-1, lines)
		if err != nil {
			log.Println(err)
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
	defer closeFile(output)

	w := bufio.NewWriter(output)
	for _, str := range lines {
		_, err = w.WriteString(str + "\n")
		if err != nil {
			log.Println(err)
		}
	}

	err = w.Flush()
	if err != nil {
		log.Println(err)
	}
}

//Получаем цифру из строки с номером нужной колонки
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
	saveMp := make(map[int][]string)
	for _, st := range lines {
		if st == "" {
			continue
		}
		num, err := getNumFromColumn(n, st)
		if err != nil {
			return nil, err
		}
		saveMp[num] = append(saveMp[num], st)
		Nums = append(Nums, num)
	}

	Result := make([]string, 0)
	sort.Ints(Nums)
	var been int
	for i, val := range Nums {
		switch {
		case been != val:
			tmpArr := saveMp[val]
			if len(tmpArr) < 2 {
				for _, st := range tmpArr {
					Result = append(Result, st)
				}
			} else {
				sortDuplicates(tmpArr, &Result, n)
			}
			been = val
		case i == 0:
			tmpArr := saveMp[val]
			sortDuplicates(tmpArr, &Result, n)
			been = val
		}

	}
	return Result, nil
}

//получаем слово из строки с номером нужной колонки
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
	column := make([]string, 0)
	saveMp := make(map[string][]string)

	for _, st := range lines {
		word := getWordFromColumn(n, st)
		if word != "" {
			word = strings.ToLower(word)
			column = append(column, word)
			saveMp[word] = append(saveMp[word], st)
		}
	}
	sort.Strings(column)
	Result := make([]string, 0)
	var been string
	for i, val := range column {
		switch {
		case been != val:
			tmpArr := saveMp[val]
			if len(tmpArr) < 2 {
				for _, st := range tmpArr {
					Result = append(Result, st)
				}
			} else {
				sortDuplicates(tmpArr, &Result, n)
			}
			been = val
		case i == 0:
			tmpArr := saveMp[val]
			sortDuplicates(tmpArr, &Result, n)
			been = val
		}

	}
	return Result
}

//Сортируем строки с одинаковыми ключами
func sortDuplicates(tmpArr []string, Result *[]string, n int) {
	//Сортируем одинаковые колонки как делает это утилита sort
	sort.Strings(tmpArr[n:])
	for _, st := range tmpArr {
		*Result = append(*Result, st)
	}
}
func deleteDuplicates(fileStrings []string, column int) []string {
	unique := make(map[string]string)
	//Ищем уникальные слова в колонке
	for _, st := range fileStrings {
		fields := strings.Fields(st)
		key := fields[column]
		if _, ok := unique[key]; !ok {
			unique[key] = st
		}
	}

	res := make([]string, 0, len(unique))
	for _, st := range unique {
		res = append(res, st)
	}

	return res
}
func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		return
	}
}
