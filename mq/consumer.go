package mq

type RabbitMQConsumer struct {
	QueueName string          // 获取接收者需要监听的队列
	RouterKey string          // 这个队列绑定的路由
	Handlers  []*HandlerIFace //消息处理器
}
