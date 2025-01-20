package main

import (
	"fmt"
	"log"

	"github.com/Lushin415/HW-otus-go/hw09_serialize/book"
	"github.com/Lushin415/HW-otus-go/hw09_serialize/serializers"
)

func main() {
	books := []book.Book{
		{ID: 1, Title: "FirstBook", Author: "Alexandr", Year: 1993, Size: 350, Rate: 4.9, Sample: []byte("Образец1")},
		{ID: 2, Title: "SecondBook", Author: "Lushin", Year: 2024, Size: 600, Rate: 2.1, Sample: []byte("Образец2")},
	}

	// JSON
	jsonData, err := serializers.SerializeToJSON(books)
	if err != nil {
		log.Fatalf("JSON serialization failed: %v", err)
	}
	fmt.Println("JSON serialized data:", string(jsonData))
	deserializedBooks, err := serializers.DeserializeFromJSON(jsonData)
	if err != nil {
		log.Fatalf("JSON deserialization failed: %v", err)
	}
	fmt.Println("Deserialized JSON:", deserializedBooks)

	// XML
	xmlData, err := serializers.SerializeToXML(books)
	if err != nil {
		log.Fatalf("XML serialization failed: %v", err)
	}
	fmt.Println("\nXML serialized data:", string(xmlData))
	deserializedBooks, err = serializers.DeserializeFromXML(xmlData)
	if err != nil {
		log.Fatalf("XML deserialization failed: %v", err)
	}
	fmt.Println("Deserialized XML:", deserializedBooks)

	// YAML
	yamlData, err := serializers.SerializeToYAML(books)
	if err != nil {
		log.Fatalf("YAML serialization failed: %v", err)
	}
	fmt.Println("\nYAML serialized data:", string(yamlData))
	deserializedBooks, err = serializers.DeserializeFromYAML(yamlData)
	if err != nil {
		log.Fatalf("YAML deserialization failed: %v", err)
	}
	fmt.Println("Deserialized YAML:", deserializedBooks)

	// Gob
	gobData, err := serializers.SerializeToGob(books)
	if err != nil {
		log.Fatalf("Gob serialization failed: %v", err)
	}
	fmt.Println("\nGob serialized data:", gobData)
	deserializedBooks, err = serializers.DeserializeFromGob(gobData)
	if err != nil {
		log.Fatalf("Gob deserialization failed: %v", err)
	}
	fmt.Println("Deserialized Gob:", deserializedBooks)

	// MessagePack
	msgpackData, err := serializers.SerializeToMessagePack(books)
	if err != nil {
		log.Fatalf("MessagePack serialization failed: %v", err)
	}
	fmt.Println("\nMessagePack serialized data:", msgpackData)
	deserializedBooks, err = serializers.DeserializeFromMessagePack(msgpackData)
	if err != nil {
		log.Fatalf("MessagePack deserialization failed: %v", err)
	}
	fmt.Println("Deserialized MessagePack:", deserializedBooks)
}
