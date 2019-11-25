package main

import (
	"com.jxtech.gather/config"
	_ "com.jxtech.gather/mq"
	"fmt"
)

func main() {
	fmt.Printf("RabbitMQ参数读取：%s\n", config.Get().RabbitMQ.User)
}
