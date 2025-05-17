package rabbitMQ

import (
	"encoding/json"
	"fmt"
	"http_server/domain"

	"github.com/streadway/amqp"
    "http_server/cmd/app/config"
)

type RabbitMQSender struct {
    connection *amqp.Connection
    channel    *amqp.Channel
    queueName  string
}

func NewRabbitMQSender(cfg config.RabbitMQ) (*RabbitMQSender, error) {
    url := fmt.Sprintf("amqp://guest:guest@%s:%d", cfg.Host, cfg.Port)
    conn, err := amqp.Dial(url)
    if err != nil {
        return nil, fmt.Errorf("connecting to rabbitMQ: %w", err)
    }

    ch, err := conn.Channel()
    if err != nil {
        return nil, err
    }

    _, err = ch.QueueDeclare(
        cfg.QueueName, // name
        true,      // durable
        false,     // delete when unused
        false,     // exclusive
        false,     // no-wait
        nil,       // arguments
    )
    if err != nil {
        return nil, err
    }

    return &RabbitMQSender{
        connection: conn,
        channel:    ch,
        queueName:  cfg.QueueName,
    }, nil
}

func (r *RabbitMQSender) Send(object domain.Object) error {
    body, err := json.Marshal(object)
    if err != nil {
        return err
    }

    err = r.channel.Publish(
        "",              // exchange
        r.queueName,     // routing key
        false,           // mandatory
        false,           // immediate
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
        })
    if err != nil {
        return err
    }

    return nil
}

func (r *RabbitMQSender) Close() {
    r.channel.Close()
    r.connection.Close()
}
