package tests

import (
	"reflect"
	"testing"

	"github.com/Lushin415/HW-otus-go/hw09_serialize/book"
	"github.com/Lushin415/HW-otus-go/hw09_serialize/serializers"
)

func TestSerializationXML(t *testing.T) {
	// Тестовые данные
	books := []book.Book{
		{ID: 1, Title: "FirstBook", Author: "Alexandr", Year: 1993, Size: 350, Rate: 4.9, Sample: []byte("Образец1")},
		{ID: 2, Title: "SecondBook", Author: "Lushin", Year: 2024, Size: 600, Rate: 2.1, Sample: []byte("Образец2")},
	}

	// XML
	t.Run("XML Serialization", func(t *testing.T) {
		data, err := serializers.SerializeToXML(books)
		if err != nil {
			t.Fatalf("XML serialization failed: %v", err)
		}
		deserializedBooks, err := serializers.DeserializeFromXML(data)
		if err != nil {
			t.Fatalf("XML deserialization failed: %v", err)
		}
		if !reflect.DeepEqual(books, deserializedBooks) {
			t.Fatalf("Deserialized XML does not match original: got %v, want %v", deserializedBooks, books)
		}
	})
}
