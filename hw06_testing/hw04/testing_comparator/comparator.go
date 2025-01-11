package comparator

import book "HW-otus/hw06_testing/hw04/testing_book"

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

// NewComparator новый компаратор со сравнением.
func NewComparator(fieldCompare PoWhat) *Comparator {
	return &Comparator{fieldCompare}
}

// Compare для сравнения книг.
func (c *Comparator) Compare(bookOne, bookTwo book.Book) bool {
	switch c.fieldCompare {
	case PoYear:
		return bookOne.Year() > bookTwo.Year()
	case PoSize:
		return bookOne.Size() > bookTwo.Size()
	case PoRate:
		return bookOne.Rate() > bookTwo.Rate()
	default:
		return false
	}
}
