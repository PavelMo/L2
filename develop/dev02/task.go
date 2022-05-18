package main

import (
	"errors"
	"fmt"
	"log"
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

		case runes[i] == '\\':
			if i+1 == len(runes) {
				i++
				continue
			}
			//Если бэкслеш, то пропускаем его и падаем дальше
			i++
			fallthrough
		case i == len(runes)-1:
			if i == len(runes)-1 {
				res.WriteRune(runes[i])
				i++
				continue
			}
			fallthrough
		case unicode.IsDigit(runes[i+1]):
			if unicode.IsDigit(runes[i+1]) {
				var (
					countSt strings.Builder
					iSkip   int
				)
				//Получаем полное число
				for j := i + 1; j < len(runes) && unicode.IsDigit(runes[j]); {
					countSt.WriteString(string(runes[j]))
					j += 1
					iSkip = j
				}
				count, err := strconv.Atoi(countSt.String())
				if err != nil {
					return "", err
				}
				res.WriteString(strings.Repeat(string(runes[i]), count))
				i = iSkip
				continue
			}
			fallthrough
		default:
			res.WriteRune(runes[i])
			i++

		}
	}
	return res.String(), nil
}
func main() {
	list := []string{"a4bc2d5e", "abcd2", "a45", "", `qwe\4\5`, `qwe\45`, `qwe\\5`, `qwe\`}

	for _, st := range list {
		unpacked, err := Unpack(st)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(unpacked)
		}

	}
}
