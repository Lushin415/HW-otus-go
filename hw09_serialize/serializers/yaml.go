package serializers

import (
	"github.com/Lushin415/HW-otus-go/hw09_serialize/book"
	"github.com/go-yaml/yaml"
)

func SerializeToYAML(books []book.Book) ([]byte, error) {
	return yaml.Marshal(books)
}

func DeserializeFromYAML(data []byte) ([]book.Book, error) {
	var books []book.Book
	err := yaml.Unmarshal(data, &books)
	return books, err
}
