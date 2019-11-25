package mq

import (
	"com.jxtech.gather/config"
	"com.jxtech.gather/gerr"
	"fmt"
	"github.com/streadway/amqp"
)

var _conn amqp.Connection

func init() {
	fmt.Println("开始RabbitMQ初始化")
	param := config.Get()
	amqpUrl := fmt.Sprintf("%s%s:%s@%s:%d/%s",
		"amqp://", param.RabbitMQ.User, param.RabbitMQ.Pass, param.RabbitMQ.Host, param.RabbitMQ.Port, param.RabbitMQ.VirtualHost)
	_conn, err := amqp.Dial(amqpUrl)
	gerr.FailOnError(err, "Failed to Dial amqp")
	defer _conn.Close()

	bindQueue(_conn)
}

func bindQueue(conn *amqp.Connection) {
	ch, err := conn.Channel()
	gerr.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	exchangeName := "exchangeCMS"
	err = ch.ExchangeDeclare(
		exchangeName,
		amqp.ExchangeTopic,
		true,
		false,
		false,
		true,
		nil,
	)
	gerr.FailOnError(err, "创建交换机失败")

	createAndBindQueue(ch, exchangeName,
		[]string{"cms.alarm.#", "cms.envdata.#"},
		[]string{"queueAlarmCMS", "queueEnvdataCMS"})

	handlerMessage(ch)
}

func createAndBindQueue(ch *amqp.Channel, exchangeName string, routes []string, queues []string) {
	for i, queueName := range queues {
		_, err := ch.QueueDeclare(
			queueName, // name
			true,      // durable
			false,     // delete when unused
			false,     // exclusive
			false,     // no-wait
			nil,       // arguments
		)

		gerr.FailOnError(err, "创建队列失败："+queueName)

		routeKey := routes[i]
		err = ch.QueueBind(queueName, routeKey, exchangeName, true, nil)
		gerr.FailOnError(err, "绑定队列失败！router:"+routeKey+";queue:"+queueName)

		fmt.Printf("成功绑定队列：router:%s --> queue:%s\n", routeKey, queueName)
	}
}

func handlerMessage(ch *amqp.Channel) {
	ch.Consume()
}
