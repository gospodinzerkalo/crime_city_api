package transport

import (
	"context"
	"encoding/json"
	amqpTr "github.com/go-kit/kit/transport/amqp"
	"github.com/gospodinzerkalo/crime_city_api/endpoint"
	"github.com/streadway/amqp"
	"github.com/go-kit/kit/log"
)

type amqpServer struct {
	sendCrime 	*amqpTr.Publisher
}

type RabbitMQConfig struct {
	Channel 	*amqp.Channel
}

func NewRabbitMqServer(endpoints endpoint.Endpoints, logg log.Logger, conf RabbitMQConfig) error {
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
				logg.Log("received \"send_crime\" request")
				qSendCrimeListener(&uppercaseDeliv)
				uppercaseDeliv.Ack(false)
			}
		}
	}()

	logg.Log("listening")
	<-forever
	return nil
}


func decodeSendCrimeAmqpRequest(_ context.Context, delivery *amqp.Delivery) (interface{}, error) {
	var req endpoint.GetCrimeRequest
	err := json.Unmarshal(delivery.Body, &req.ID)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeSendCrimeAmqpResponse(context.Context, *amqp.Publishing, interface{}) error {
	return nil
}