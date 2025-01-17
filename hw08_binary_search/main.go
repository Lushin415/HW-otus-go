package main

import (
	"fmt"

	"github.com/Lushin415/HW-otus-go/hw08_binary_search/binary"
)

func main() {
	fmt.Println("Введите 2 числа, первое начало массива, второе - его конец")
	var a, b, c int
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		return
	}
	items := binary.CreateArray(a, b)
	fmt.Println("Массив готов, напишите число для поиска")
	_, err = fmt.Scan(&c)
	if err != nil {
		return
	}
	fmt.Println(binary.SearchBinary(items, c))
}
