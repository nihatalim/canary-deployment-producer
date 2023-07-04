package main

import (
	"context"
	"fmt"
	"producer/config"
	"strings"
	"time"

	"github.com/Trendyol/kafka-konsumer"
)

func main() {
	ctx := context.Background()
	cfg := config.GetConfig()

	go startProduce(ctx, cfg)

	fmt.Println("Producer started...")

	<-ctx.Done()

	fmt.Println("Producer stopped...")
}

func startProduce(ctx context.Context, cfg *config.AppConfig) {
	t := time.NewTicker(time.Second)

	p, _ := kafka.NewProducer(kafka.ProducerConfig{
		ClientID: cfg.Kafka.ClientId,
		Writer: kafka.WriterConfig{
			Brokers: strings.Split(cfg.Kafka.Brokers, ","),
		},
	})

	for {
		select {
		case <-ctx.Done():
			return

		case tick := <-t.C:
			currentTime := tick.String()

			msg := fmt.Sprintf("{\"time\": \"%s\"}", currentTime)

			_ = p.Produce(context.Background(), kafka.Message{
				Topic: cfg.Kafka.Topic,
				Key:   []byte(currentTime),
				Value: []byte(msg),
			})

			fmt.Println("Message produced:" + msg)
		}
	}
}
