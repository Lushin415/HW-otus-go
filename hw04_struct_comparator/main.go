package main

import "fmt"

// Book Структура с неэкспортируемыми полями.
type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}

// GetID Методы для получения полей структуры Book.
func (b *Book) GetID() int {
	return b.id
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) GetAuthor() string {
	return b.author
}

func (b *Book) GetYear() int {
	return b.year
}

func (b *Book) GetSize() int {
	return b.size
}

func (b *Book) GetRate() float64 {
	return b.rate
}

// PoWhat сравнить "По" (году, размеру, рейтингу).
type PoWhat int

const (
	PoYear PoWhat = iota
	PoSize
	PoRate
)

// Comparator структура для хранения сравнения.
type Comparator struct {
	fieldCompare PoWhat
}

// NewComparator новый компаратор с сравнением.
func NewComparator(fieldCompare PoWhat) *Comparator {
	return &Comparator{fieldCompare}
}

// Compare для сравнения книг.
func (c *Comparator) Compare(bookOne, bookTwo *Book) bool {
	switch c.fieldCompare {
	case PoYear:
		return bookOne.GetYear() > bookTwo.GetYear()
	case PoSize:
		return bookOne.GetSize() > bookTwo.GetSize()
	case PoRate:
		return bookOne.GetRate() > bookTwo.GetRate()
	default:
		return false
	}
}

func main() {
	book1 := &Book{id: 1, title: "FirstBook", author: "Alexandr", year: 1993, size: 350, rate: 4.9}
	book2 := &Book{id: 2, title: "SecondBook", author: "Lushin", year: 2024, size: 600, rate: 2.1}
	// Создаем компаратор для сравнения по году
	yCompare := NewComparator(PoYear)
	fmt.Println("Книга № 1 больше чем Книга 2 по году:", yCompare.Compare(book1, book2))
	sCompare := NewComparator(PoSize)
	fmt.Println("Книга № 1 больше чем Книга 2 по размеру:", sCompare.Compare(book1, book2))
	rCompare := NewComparator(PoRate)
	fmt.Println("Книга № 1 больше чем Книга 2 по рейтингу:", rCompare.Compare(book1, book2))
}
