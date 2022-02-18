package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic          = "supplier-product-count-update"
	broker1Address = "localhost:9092"
)

type KafkaMsg struct {
	Sid string `json:"sid"`
	Cmd string `json:"cmd"`
}

func (s KafkaMsg) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s KafkaMsg) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &s)
}

func Produce(ctx context.Context, givenSid string, givenCmd string) {

	key := 0
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      []string{broker1Address},
		Topic:        topic,
		BatchSize:    2,
		BatchTimeout: 2 * time.Second,
	})
	m := KafkaMsg{givenSid, givenCmd}
	b, err := json.Marshal(m)
	if err != nil {
		panic("Error " + err.Error())
	}
	err = w.WriteMessages(ctx, kafka.Message{
		Key:   []byte(strconv.Itoa(key)),
		Value: b,
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}

	fmt.Println("writes:", givenSid, " ", givenCmd)

	key++
}

func Consume(ctx context.Context) {
	fmt.Println("consumer starts..")

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{broker1Address},
		Topic:       topic,
		GroupID:     "my-group",
		MinBytes:    5,
		MaxBytes:    1e6,
		MaxWait:     3 * time.Second,
		StartOffset: kafka.LastOffset,
	})
	msg, err := r.ReadMessage(ctx)
	if err != nil {
		panic("could not read message " + err.Error())
	}

	var kmsg KafkaMsg
	parseErr := json.Unmarshal(msg.Value, &kmsg)
	if parseErr != nil {
		panic(parseErr)
	}

	fmt.Print("received: ", kmsg.Sid, " ")
	if kmsg.Cmd == "DEL" {
		_, err = http.Get("http://localhost:3000/api/updateSupplierProductCountDEC/" + string(kmsg.Sid))
		if err != nil {
			log.Fatal(err)
		}

	}
	if kmsg.Cmd == "ADD" {
		_, err = http.Get("http://localhost:3000/api/updateSupplierProductCountINC/" + string(kmsg.Sid))
		if err != nil {
			log.Fatal(err)
		}

	}
	fmt.Print(kmsg.Cmd)
	fmt.Println()

}
