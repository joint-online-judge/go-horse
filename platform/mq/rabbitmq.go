package mq

import (
	"fmt"

	"github.com/joint-online-judge/go-horse/pkg/configs"
	"github.com/joint-online-judge/go-horse/pkg/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

var Conn *amqp.Connection

func ConnectRabbitMQ() {
	conf := configs.Conf
	config := amqp.Config{
		Vhost:      conf.RabbitmqVhost,
		Properties: amqp.NewConnectionProperties(),
	}
	config.Properties.SetClientConnectionName("go-horse")
	uri := fmt.Sprintf(
		"amqp://%s:%s@%s:%d/",
		conf.RabbitmqUsername,
		conf.RabbitmqPassword,
		conf.RabbitmqHost,
		conf.RabbitmqPort,
	)
	Conn, err := amqp.DialConfig(uri, config)
	if err != nil {
		logger.Fatalf("failed to dial to rabbitmq: %+v", err)
	}
	channel, err := Conn.Channel()
	if err != nil {
		logger.Fatalf("fail to get a channel in rabbitmq: %+v", err)
	}
	defer channel.Close()
}
