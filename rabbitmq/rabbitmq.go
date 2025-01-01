package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/RishangS/customerror"
	"github.com/RishangS/env"

	"github.com/rabbitmq/amqp091-go"
)

var (
	rmqURL      = env.GetEnv("RMQ_URL", "localhost:5672")
	rmqUser     = env.GetEnv("RMQ_USER", "guest")
	rmqPassword = env.GetEnv("RMQ_PASSWORD", "guest")
)

type RabbitMQ struct {
	connection *amqp091.Connection
	channel    *amqp091.Channel
	queue      map[string]amqp091.Queue
	mu         sync.Mutex
}

func NewRabbitMQService() *RabbitMQ {
	connection, err := amqp091.Dial(fmt.Sprintf("amqp://%s:%s@%s/", rmqUser, rmqPassword, rmqURL))
	if err != nil {
		log.Fatalf("RabbitMQ connection failed %s", err)
	}
	ch, err := connection.Channel()
	if err != nil {
		log.Fatalf("RabbitMQ create channel failed %s", err)
	}

	return &RabbitMQ{connection: connection, channel: ch, queue: make(map[string]amqp091.Queue)}
}

func (r *RabbitMQ) CreateQueue(queueName string) error {
	var durable, autoDelete, exclusive, noWait = true, false, false, false
	queue, err := r.channel.QueueDeclare(queueName, durable, autoDelete, exclusive, noWait, nil)
	if err != nil {
		return customerror.NewQueueCreationErrorExceptionError(fmt.Sprintf("%s error: %v", queueName, err.Error()))
	}
	r.addNewQueue(queueName, queue)
	return nil
}

func (r *RabbitMQ) addNewQueue(queueName string, queue amqp091.Queue) {
	r.queue[queueName] = queue
}

func (r *RabbitMQ) IsConnectionAlive() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.connection == nil || r.connection.IsClosed() {
		return false
	}
	closeChan := r.connection.NotifyClose(make(chan *amqp091.Error, 1))

	select {
	case err := <-closeChan:
		if err != nil {
			return false
		}
	default:
		log.Println("rabbitmq connection alive")
	}
	return true
}

func (r *RabbitMQ) CheckAndReconnect() error {
	if r.IsConnectionAlive() {
		return nil
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	if r.connection != nil && !r.connection.IsClosed() {
		_ = r.connection.Close()
	}
	connection, err := amqp091.Dial(fmt.Sprintf("amqp://%s:%s@%s/", rmqUser, rmqPassword, rmqURL))
	if err != nil {
		log.Fatalf("RabbitMQ connection failed %s", err)
	}
	channel, err := connection.Channel()
	if err != nil {
		log.Fatalf("RabbitMQ create channel failed %s", err)
	}

	//FIXME on reconnection if queues are lost then loop on all queues and recreate them
	r.connection = connection
	r.channel = channel
	return nil
}

func (r *RabbitMQ) GetQueue(queueName string) *amqp091.Queue {
	q := r.queue[queueName]
	return &q
}

func (r *RabbitMQ) Publish(queueName string, message interface{}) error {
	var mandatory, immediate = false, false
	var exchange = ""

	bytes, err := json.Marshal(message)
	if err != nil {
		return customerror.NewMarshalExceptionError(err.Error())
	}

	publishPayload := amqp091.Publishing{
		ContentType: "text/plain",
		Body:        bytes,
	}

	return r.channel.Publish(exchange, queueName, mandatory, immediate, publishPayload)
}

func (r *RabbitMQ) Consume(queueName string) (<-chan amqp091.Delivery, error) {
	var autoAck, exclusive, noLocal, noWait = true, false, false, false
	consumerTag := ""
	msgsChan, err := r.channel.Consume(queueName, consumerTag, autoAck, exclusive, noLocal, noWait, nil)
	if err != nil {
		return nil, customerror.NewConsumerRegisterExceptionError(err.Error())
	}
	return msgsChan, nil
}
