package tests

import (
	"github.com/Lushin415/HW-otus-go/hw09_serialize/serializers"
	"reflect"
	"testing"

	"github.com/Lushin415/HW-otus-go/hw09_serialize/book"
)

func TestSerialization(t *testing.T) {
	// Тестовые данные
	books := []book.Book{
		{ID: 1, Title: "FirstBook", Author: "Alexandr", Year: 1993, Size: 350, Rate: 4.9, Sample: []byte("Образец1")},
		{ID: 2, Title: "SecondBook", Author: "Lushin", Year: 2024, Size: 600, Rate: 2.1, Sample: []byte("Образец2")},
	}

	// JSON
	t.Run("JSON Serialization", func(t *testing.T) {
		data, err := serializers.SerializeToJSON(books)
		if err != nil {
			t.Fatalf("JSON serialization failed: %v", err)
		}
		deserializedBooks, err := serializers.DeserializeFromJSON(data)
		if err != nil {
			t.Fatalf("JSON deserialization failed: %v", err)
		}
		if !reflect.DeepEqual(books, deserializedBooks) {
			t.Fatalf("Deserialized JSON does not match original: got %v, want %v", deserializedBooks, books)
		}
	})

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

	// Gob
	t.Run("Gob Serialization", func(t *testing.T) {
		data, err := serializers.SerializeToGob(books)
		if err != nil {
			t.Fatalf("Gob serialization failed: %v", err)
		}
		deserializedBooks, err := serializers.DeserializeFromGob(data)
		if err != nil {
			t.Fatalf("Gob deserialization failed: %v", err)
		}
		if !reflect.DeepEqual(books, deserializedBooks) {
			t.Fatalf("Deserialized Gob does not match original: got %v, want %v", deserializedBooks, books)
		}
	})

	// MessagePack
	t.Run("MessagePack Serialization", func(t *testing.T) {
		data, err := serializers.SerializeToMessagePack(books)
		if err != nil {
			t.Fatalf("MessagePack serialization failed: %v", err)
		}
		deserializedBooks, err := serializers.DeserializeFromMessagePack(data)
		if err != nil {
			t.Fatalf("MessagePack deserialization failed: %v", err)
		}
		if !reflect.DeepEqual(books, deserializedBooks) {
			t.Fatalf("Deserialized MessagePack does not match original: got %v, want %v", deserializedBooks, books)
		}
	})
}
