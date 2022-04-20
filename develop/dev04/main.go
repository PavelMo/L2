package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===
Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.
Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func findAnagrams(dictionary []string) map[string][]string {
	res := make(map[string][]string)
	unique := make(map[string]struct{})

	for _, val := range dictionary {

		lowerWord := strings.ToLower(val)
		sortedChars := getChars(lowerWord)

		if _, ok := unique[lowerWord]; !ok {
			res[sortedChars] = append(res[sortedChars], lowerWord)
			unique[lowerWord] = struct{}{}
		}
	}

	return res
}
func getChars(st string) string {
	chars := strings.Split(st, "")
	sort.Strings(chars)

	return strings.Join(chars, "")
}

func main() {

	dictionary := []string{
		"Тяпка",
		"Пятак",
		"Пятак",
		"пятка",
		"слиток",
		"слиток",
		"столик",
		"листок",
		"Топот",
		"Потоп",
	}

	anagrams := findAnagrams(dictionary)

	for key, value := range anagrams {
		if len(value) < 2 {
			delete(anagrams, key)
		} else {
			fmt.Printf("Key: %s\nValue: %v\n\n", key, value)
		}

	}
}

