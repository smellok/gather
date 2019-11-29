package mq

import (
	"fmt"
)

type RabbitMQ struct {
	Host         string
	Port         int
	User         string
	Pass         string
	VHost        string
	ExchangeName string                         //交换机名称
	BindingMap   map[string]string              //主题路由绑定
	consumers    map[string][]*RabbitMQConsumer //主题消费者
}

func (mq *RabbitMQ) AddConsumer(topic string, consumer *RabbitMQConsumer) {
	consumerSlice := mq.consumers[topic]
	if nil == consumerSlice {
		consumerSlice = []*RabbitMQConsumer{consumer}
	} else {
		exist := false
		for _, existConsumer := range consumerSlice {
			if existConsumer == consumer {
				fmt.Printf("当前主题：%s,消费者：%v已经注册过了！", topic, consumer)
				exist = true
				break
			}
		}

		if !exist {
			consumerSlice = append(consumerSlice, consumer)
		}
	}
}

func (mq *RabbitMQ) RmConsumer(topic string, consumer *RabbitMQConsumer) {
	consumerSlice := mq.consumers[topic]
	if nil == consumerSlice {
		return
	}

	idx := -1
	for i, existConsumer := range consumerSlice {
		if existConsumer == consumer {
			fmt.Printf("当前主题：%s,消费者：%v找到了！", topic, consumer)
			idx = i
			break
		}
	}

	if idx >= 0 {
		consumerSlice = append(consumerSlice[:idx], consumerSlice[idx+1:]...)
	}
}
