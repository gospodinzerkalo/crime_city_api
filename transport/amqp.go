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

type RabbitMQConfig struct {
	Channel 	*amqp.Channel
}

func NewRabbitMqServer(endpoints endpoint.Endpoints, log *logger.Logger, conf RabbitMQConfig) error {
	qSendCrime, err := conf.Channel.QueueDeclare(
		"send_crime", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return err
	}

	qSendCrimeMsgs, err := conf.Channel.Consume(
		qSendCrime.Name,
		"",    // consumer
		false, // autoAck
		false, // exclusive
		false, // noLocal
		false, // noWait
		nil,   // args
	)

	qSendCrimeAMQPHandler := amqpTr.NewSubscriber(
		endpoints.GetCrime,
		decodeSendCrimeAmqpRequest,
		encodeSendCrimeAmqpResponse,
	)

	qSendCrimeListener := qSendCrimeAMQPHandler.ServeDelivery(conf.Channel)

	forever := make(chan bool)

	go func() {
		for true {
			select {
			case uppercaseDeliv := <-qSendCrimeMsgs:
				log.Noticef("received \"send_crime\" request")
				qSendCrimeListener(&uppercaseDeliv)
				uppercaseDeliv.Ack(false)
			}
		}
	}()

	log.Noticef("listening")
	<-forever
	return nil
}


func decodeSendCrimeAmqpRequest(context.Context, *amqp.Delivery) (request interface{}, err error) {
	return nil, nil
}

func encodeSendCrimeAmqpResponse(context.Context, *amqp.Publishing, interface{}) error {
	return nil
}