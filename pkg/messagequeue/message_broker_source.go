package messagequeue

import (
	"github.com/rabbitmq/amqp091-go"
	"log"
)

func Init() (*amqp091.Channel, error) {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatalln(err)
	}

	return conn.Channel()
}
