package serializers

import (
	"encoding/xml"

	"github.com/Lushin415/HW-otus-go/hw09_serialize/book"
)

func SerializeToXML(books []book.Book) ([]byte, error) {
	return xml.Marshal(book.List{Books: books})
}

func DeserializeFromXML(data []byte) ([]book.Book, error) {
	var bookList book.List
	err := xml.Unmarshal(data, &bookList)
	return bookList.Books, err
}
