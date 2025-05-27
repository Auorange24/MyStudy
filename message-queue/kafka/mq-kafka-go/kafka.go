package mq_kafka_go

import (
	"context"
	"log"
	"github.com/segmentio/kafka-go"
)

func fatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func WriteByConn() {
	topic := "my-topic" // 主题
	partition := 0      // 分区

	// 连接到Kafka的集群节点
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	fatalError(err)

	// 设置发送消息的超时时间
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("Hello, World!")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func ReadByConn() {
	topic := "my-topic"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	fatalError(err)

	
}