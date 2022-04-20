package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===
Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)
В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.
Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
var errString = errors.New("incorrect input")

func Unpack(st string) (string, error) {
	var res strings.Builder

	if len(st) < 1 || unicode.IsDigit(rune(st[0])) {
		return "", errString
	}

	runes := []rune(st)

	for i := 0; i < len(runes); {
		switch {
		case i == len(runes)-1:
			res.WriteRune(runes[i])
			i++

		case unicode.IsDigit(runes[i+1]):
			var (
				countSt strings.Builder
				iSkip   int
			)

			for j := i + 1; j < len(runes) && unicode.IsDigit(runes[j]); {
				countSt.WriteString(string(runes[j]))
				j += 1
				iSkip = j
			}
			count, err := strconv.Atoi(countSt.String())
			if err != nil {
				fmt.Println(err)
			}
			res.WriteString(strings.Repeat(string(runes[i]), count))
			i = iSkip
		default:
			res.WriteRune(runes[i])
			i++

		}
	}
	return res.String(), nil
}
func main() {
	strs := []string{"a4bc2d5e", "abcd2", "45", ""}

	for _, st := range strs {
		unpacked, err := Unpack(st)
		if err != nil {
			_, err = fmt.Fprintln(os.Stderr, err)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(unpacked)
		}

	}
}
