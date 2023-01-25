package mq

import (
	"fmt"

	"github.com/joint-online-judge/go-horse/pkg/configs"
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
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
		log.Fatalf("failed to dial to rabbitmq: %+v", err)
	}
	channel, err := Conn.Channel()
	if err != nil {
		log.Fatalf("fail to get a channel in rabbitmq: %+v", err)
	}
	defer channel.Close()
}
