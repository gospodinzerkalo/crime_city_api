package rabbitmq

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	amqpTransport "github.com/go-kit/kit/transport/amqp"
	"log"
	"sync"
)

const (
	SENDCRIME = "send_crime"
)

type Config struct {
	URL 		string
}

type rabbitMqInterface interface {
	SendCrimeMsg(crime Crime) error
}

type rabbitMq struct {
	conn *amqp.Connection
}

func NewRabbitMQPublisher(config Config) (rabbitMqInterface, error){
	conn, err := amqp.Dial(config.URL)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	return rabbitMq{conn: conn}, nil
}

func (r rabbitMq) SendCrimeMsg(crime Crime) error {
	ch, err := r.conn.Channel()
	if err != nil {
		return err
	}
	replyQueue, err := ch.QueueDeclare(
		SENDCRIME,
		false, // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,   //args
	)

	sendCrimePuplisher := amqpTransport.NewPublisher(
		ch,
		&replyQueue,
		encodeSendCrimeAMQPRequest,
		decodeSendCrimeAMQPResponse,
		amqpTransport.PublisherBefore(
			// queue name specified by subscriber
			amqpTransport.SetPublishKey("send_crime"),
		),
	)

	var wg sync.WaitGroup
	uppercaseRequest := stringsvc.UppercaseRequest{S: *word}
	wg.Add(1)
	go func() {
		log.Println("sending uppercase request")
		defer wg.Done()
		response, err := uppercaseEndpoint(context.Background(), uppercaseRequest)
		if err != nil {
			log.Println("error with uppercase request", err)
		} else {
			log.Println("uppercase response: ", response)
		}
	}()

	return nil
}

func encodeSendCrimeAMQPRequest(ctx context.Context, publishing *amqp.Publishing, request interface{}) error {
	uppercaseRequest, ok := request.(stringsvc.UppercaseRequest)
	if !ok {
		return badRequest
	}
	b, err := json.Marshal(uppercaseRequest)
	if err != nil {
		return err
	}
	publishing.Body = b
	return nil
}

func decodeSendCrimeAMQPResponse(ctx context.Context, delivery *amqp.Delivery) (interface{}, error) {
	var response stringsvc.UppercaseResponse
	err := json.Unmarshal(delivery.Body, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}