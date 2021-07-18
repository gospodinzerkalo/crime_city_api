package transport

import (
	"context"
	amqpTr "github.com/go-kit/kit/transport/amqp"
	"github.com/gospodinzerkalo/crime_city_api/endpoint"
	"github.com/nats-io/nats-server/v2/logger"
	"github.com/streadway/amqp"
	"log"
)

type amqpServer struct {
	sendCrime 	*amqpTr.Publisher
}

type RabbitMQConfig struct {
	Channel 	*amqp.Channel
	Queue 		*amqp.Queue
}

func NewRabbitMqServer(endpoints endpoint.Endpoints, log *logger.Logger, conf RabbitMQConfig) *amqpServer {
	return &amqpServer{
		sendCrime: amqpTr.NewPublisher(
			conf.Channel,
			conf.Queue,
			encodeSendCrimeAmqpResponse,
			decodeSendCrimeAmqpRequest,
			),
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func decodeSendCrimeAmqpRequest(context.Context, *amqp.Delivery) (request interface{}, err error) {
	return nil, nil
}

func encodeSendCrimeAmqpResponse(context.Context, *amqp.Publishing, interface{}) error {
	return nil
}