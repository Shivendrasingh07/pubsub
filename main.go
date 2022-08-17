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

	go NewKafkaProvider(ctx)
	//Subscribe()

	Subscribe2(TopicRealtimeMessage, ctx)

}

func NewKafkaProvider(ctx context.Context) {

	fmt.Print("Enter text: ")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
	}

	outboundMessageBytes, err := json.Marshal(&Message{
		Message: input + string(TopicChatMessage),
	})

	if err != nil {
		logrus.Errorf("toggleBlock: error marshal outbound message data  %v", err)

	}
	kafkaHost := "127.0.0.1:9092"
	// TopicChatMessage
	chatWriter := &kafka.Writer{
		Addr:      kafka.TCP(kafkaHost),
		Topic:     string(TopicChatMessage),
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

	outboundMessageBytes, err = json.Marshal(&Message{
		Message: input + string(TopicRealtimeMessage),
	})

	//TopicRealtimeMessage
	chatWriterRealtime := &kafka.Writer{
		Addr:      kafka.TCP(kafkaHost),
		Topic:     string(TopicRealtimeMessage),
		BatchSize: 1,
	}
	err = chatWriterRealtime.WriteMessages(context.Background(),
		kafka.Message{
			Value: outboundMessageBytes,
		},
	)
	if err != nil {
		logrus.Errorf("Publish: failed to write chat messages: %v", err)
	}

}

func Subscribe2(topic Topic, ctx context.Context) {
	kafkaHost := "127.0.0.1:9092"

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaHost},
		Topic:   string(topic),
		GroupID: "g1",
	})

	writeMessage(ctx, r)
}

func writeMessage(ctx context.Context, r *kafka.Reader) {
	for {
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
