package comparator

// Book Структура с неэкспортируемыми полями.
type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}

// ID геттеры для получения полей структуры.
func (b *Book) ID() int {
	return b.id
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) Year() int {
	return b.year
}

func (b *Book) Size() int {
	return b.size
}

func (b *Book) Rate() float64 {
	return b.rate
}

// Сеттеры для получения полей структуры Book.

func (b *Book) SetID(id int) {
	b.id = id
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b *Book) SetRate(rate float64) {
	b.rate = rate
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
		return bookOne.Year() > bookTwo.Year()
	case PoSize:
		return bookOne.Size() > bookTwo.Size()
	case PoRate:
		return bookOne.Rate() > bookTwo.Rate()
	default:
		return false
	}
}
