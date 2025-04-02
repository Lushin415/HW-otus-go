package serializers

import (
	"github.com/Lushin415/HW-otus-go/hw09_serialize/book"
	"github.com/vmihailenco/msgpack/v5"
)

func SerializeToMessagePack(books []book.Book) ([]byte, error) {
	return msgpack.Marshal(books)
}

func DeserializeFromMessagePack(data []byte) ([]book.Book, error) {
	var books []book.Book
	err := msgpack.Unmarshal(data, &books)
	return books, err
}
