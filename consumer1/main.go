package main

import (
	"context"
	"fmt"
	"log"

	kafka "github.com/segmentio/kafka-go"
)

func main() {

	topic := "example-123"

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		GroupID: "group-1",
		Topic:   topic,
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Println("Reading from Consumer 1 with GroupID 1")
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

// install and run kafka locally.
// write a producer (write to a topic "example-123").
// write a consumer (read from a topic "example-123",).
// 3 consumers -> 2 of them in the same consumer group, 1 of them in a separate consumer group.
