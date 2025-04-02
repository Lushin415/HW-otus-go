package tests

import (
	"reflect"
	"testing"

	"github.com/Lushin415/HW-otus-go/hw09_serialize/book"
	"github.com/Lushin415/HW-otus-go/hw09_serialize/serializers"
)

func TestSerializationYAML(t *testing.T) {
	// Тестовые данные
	books := []book.Book{
		{ID: 1, Title: "FirstBook", Author: "Alexandr", Year: 1993, Size: 350, Rate: 4.9, Sample: []byte("Образец1")},
		{ID: 2, Title: "SecondBook", Author: "Lushin", Year: 2024, Size: 600, Rate: 2.1, Sample: []byte("Образец2")},
	}

	// YAML
	t.Run("YAML Serialization", func(t *testing.T) {
		data, err := serializers.SerializeToYAML(books)
		if err != nil {
			t.Fatalf("YAML serialization failed: %v", err)
		}
		deserializedBooks, err := serializers.DeserializeFromYAML(data)
		if err != nil {
			t.Fatalf("YAML deserialization failed: %v", err)
		}
		if !reflect.DeepEqual(books, deserializedBooks) {
			t.Fatalf("Deserialized YAML does not match original: got %v, want %v", deserializedBooks, books)
		}
	})
}
