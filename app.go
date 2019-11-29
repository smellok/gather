package main

import (
	"github.com/smellok/gather/handler"
	"github.com/smellok/gather/mq"
)

func main() {

	var deviceHandler mq.HandlerIFace = &handler.CmsDeviceHandler{}

	deviceConsumer := &mq.RabbitMQConsumer{
		QueueName: "queueAlarmCMS",
		RouterKey: "cms.alarm.#",
		Handlers:  []*mq.HandlerIFace{&deviceHandler},
	}
	print(deviceConsumer)
}
