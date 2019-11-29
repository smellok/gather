package mq

type RabbitMQConsumer struct {
	Name    string
	Handler HandlerIFace //消息处理器
}
