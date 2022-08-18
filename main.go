package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"os"
)

type Topic string

const (
	TopicChatMessage     Topic = "chat_message"
	TopicRealtimeMessage Topic = "real_time_events_message"
)

type Message struct {
	ToUserID string
	Message  string
}

type KafkaWriter struct {
	ChatWriterChat     *kafka.Writer
	ChatWriterRealTime *kafka.Writer
}

func main() {

	fmt.Println("start main")
	ctx := context.Background()

	go publish(ctx)
	//Subscribe()

	subscribe(TopicChatMessage, ctx)

}

func publish(ctx context.Context) {

	fmt.Print("Enter text: ")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
	}

	writeMessage(TopicChatMessage, input, ctx)

}

func writeMessage(topic2 Topic, input string, ctx context.Context) {
	outboundMessageBytes, err := json.Marshal(&Message{
		Message: input + string(topic2),
	})

	if err != nil {
		logrus.Errorf("toggleBlock: error marshal outbound message data  %v", err)
	}

	kafkaHost := "127.0.0.1:9092"
	// TopicChatMessage
	chatWriter := &kafka.Writer{
		Addr:      kafka.TCP(kafkaHost),
		Topic:     string(topic2),
		BatchSize: 1,
	}

	err = chatWriter.WriteMessages(ctx,
		kafka.Message{
			Value: outboundMessageBytes,
		},
	)
	if err != nil {
		logrus.Errorf("Publish: failed to write chat messages: %v", err)
	}

	if err = chatWriter.Close(); err != nil {
		logrus.Errorf("Publish: failed to chatWriter.Close() : %v", err)
	}
}

func subscribe(topic Topic, ctx context.Context) {
	kafkaHost := "127.0.0.1:9092"

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaHost},
		Topic:   string(topic),
		GroupID: "g1",
	})

	readMessage(ctx, r)
}

func readMessage(ctx context.Context, r *kafka.Reader) {

	m, err := r.ReadMessage(ctx)
	if err != nil {
		logrus.Errorf("Failed to Read Message from topic: %q error: %v", m.Topic, err)
	}

	var message Message
	if err = json.Unmarshal(m.Value, &message); err != nil {
		logrus.Errorf("Failed to Unmarshal message from topic: %q error: %v", m.Topic, err)
		return
	}
	fmt.Println(message)
}

//func Subscribe() {
//	kafkaHost := "127.0.0.1:9092"
//
//	r := kafka.NewReader(kafka.ReaderConfig{
//		Brokers: []string{kafkaHost},
//		Topic:   string(TopicRealtimeMessage),
//		GroupID: "g1",
//	})
//	//_ = r.SetOffset(kafka.LastOffset)
//
//	m, err := r.ReadMessage(context.Background())
//	if err != nil {
//		logrus.Errorf("Failed to Read Message from topic: %q error: %v", m.Topic, err)
//	}
//
//	var message Message
//	if err = json.Unmarshal(m.Value, &message); err != nil {
//		logrus.Errorf("Failed to Unmarshal message from topic: %q error: %v", m.Topic, err)
//		return
//	}
//
//	//logrus.Println("real_time_events_message")
//	fmt.Println(message)
//}
