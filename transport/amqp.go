package transport

import (
	"context"
	amqpTr "github.com/go-kit/kit/transport/amqp"
	"github.com/gospodinzerkalo/crime_city_api/endpoint"
	"github.com/nats-io/nats-server/v2/logger"
	"github.com/streadway/amqp"
)

type amqpServer struct {
	sendCrime 	*amqpTr.Publisher
}

func NewRabbitMqServer(endpoints endpoint.Endpoints, log logger.Logger) *amqpServer {
	return &amqpServer{
		sendCrime: amqpTr.NewPublisher(nil,nil,encodeSendCrimeAmqpResponse, decodeSendCrimeAmqpRequest),
	}
}

func decodeSendCrimeAmqpRequest(context.Context, *amqp.Delivery) (request interface{}, err error) {
	return nil, nil
}

func encodeSendCrimeAmqpResponse(context.Context, *amqp.Publishing, interface{}) error {
	return nil
}