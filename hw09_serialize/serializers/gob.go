package serializers

import (
	"bytes"
	"encoding/gob"

	"github.com/Lushin415/HW-otus-go/hw09_serialize/book"
)

func SerializeToGob(books []book.Book) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(books)
	return buffer.Bytes(), err
}

func DeserializeFromGob(data []byte) ([]book.Book, error) {
	var books []book.Book
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	err := decoder.Decode(&books)
	return books, err
}
