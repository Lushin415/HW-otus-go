package main

import "fmt"

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}

const (
	idField = iota
	titleField
	authorField
	yearField
	sizeField
	rateField
)

type Field int

func (b *Book) CompareField(o Book, f Field) bool {
	switch f {
	case idField:
		return b.id > o.id
	case titleField:
		return b.title > o.title
	case authorField:
		return b.author > o.author
	case yearField:
		return b.year > o.year
	case sizeField:
		return b.size > o.size
	case rateField:
		return b.rate > o.rate
	default:
		return false
	}
}
func (b *Book) GetBook() string {
	return fmt.Sprintf("ID: %d, Title: %s, Author: %s, Year: %d, Size: %d, Rate: %.1f", b.id, b.title, b.author, b.year, b.size, b.rate)
}

func main() {
	b1 := Book{
		id:     100,
		title:  "First",
		author: "Sasha",
		year:   2024,
		size:   55,
		rate:   4.8}
	details := b1.GetBook()
	fmt.Println(details)

	b2 := Book{
		id:     150,
		title:  "Second",
		author: "Alexandr",
		year:   2023,
		size:   45,
		rate:   3.2}
	details = b2.GetBook()
	fmt.Println(details)

	fields := []Field{idField, titleField, authorField, yearField, sizeField, rateField}
	for _, f := range fields {
		result := b1.CompareField(b2, f)
		fmt.Printf("Сравнение книги 1 и книги 2 %d: %v\n", f, result)
	}

}
