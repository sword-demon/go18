package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"net"
	"strconv"
	"testing"
)

func TestListTopics(t *testing.T) {
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}

	for k := range m {
		fmt.Println(k)
	}
}

func TestCreateTopic(t *testing.T) {
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}
	topic := "audit_go18"
	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     3,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWriteMessage(t *testing.T) {
	publisher := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "audit_go18",
		Balancer: &kafka.LeastBytes{},
		// AllowAutoTopicCreation the topic will be created automatically if it does not exist
		AllowAutoTopicCreation: true,
	}
	defer publisher.Close()

	err := publisher.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Hello world!"),
		},
		kafka.Message{
			Key:   []byte("Key-B"),
			Value: []byte("Hello kafka!"),
		},
		kafka.Message{
			Key:   []byte("Key-C"),
			Value: []byte("Hello go!"),
		},
	)

	if err != nil {
		t.Fatalf("failed to write messages: %v", err)
	} else {
		fmt.Println("Messages written successfully")
	}
}

func TestReadMessage(t *testing.T) {
	subscriber := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  "devcloud-go18-audit",
		Topic:    "audit_go18",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	defer subscriber.Close()

	for {
		m, err := subscriber.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

		// 处理完消息后需要提交该消息已经消费完成,消费者挂掉之后保存消息消费的状态
		if err := subscriber.CommitMessages(context.Background(), m); err != nil {
			t.Fatal("failed to commit messages: ", err)
		}
	}
}
