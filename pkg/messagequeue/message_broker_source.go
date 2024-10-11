package messagequeue

import (
	"github.com/rabbitmq/amqp091-go"
	"log/slog"
	"time"
)

func Init() (*amqp091.Channel, error) {
	conn, err := amqp091.Dial("amqp://guest:guest@rabbitmq:5672/")

	if err != nil {
		time.Sleep(5 * time.Second)
		slog.Warn("error", err)
		return Init()
	}

	channel, err := conn.Channel()

	if err != nil {
		return channel, nil
	}

	_, err = channel.QueueDeclare("transfer-money", true, false, false, false, nil)
	_, _ = channel.QueueDeclare("transfer-money-dlq", true, false, false, false, nil)

	if err != nil {
		return channel, err
	}

	return channel, nil
}
