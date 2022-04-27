Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передачи их в качестве аргументов функции.

```go
package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
```

Ответ:
```
Срез состоит из указателя на массив, длины и ёмкости (максимальная длина).

При передачи в функцию, можно менять уже существующие элементы внутри с
лайса, так как меняются они в массиве, но если происходит выделение памяти, 
и нижлежащий массив меняется, то мы меняем уже в новом массиве, 
на который ссылка пропадает после выхода из функции

```