package serializers

import (
	"encoding/json"

	"github.com/Lushin415/HW-otus-go/hw09_serialize/book"
)

func SerializeToJSON(books []book.Book) ([]byte, error) {
	return json.Marshal(books)
}

func DeserializeFromJSON(data []byte) ([]book.Book, error) {
	var books []book.Book
	err := json.Unmarshal(data, &books)
	return books, err
}
