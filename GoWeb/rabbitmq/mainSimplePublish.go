package main

import (
	"Interview_Questions/GoWeb/rabbitmq/RabbitMQ"
	"fmt"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("qixiaSimple")
	rabbitmq.PublishSimple("Hello qixia")
	fmt.Println("发送成功！")
}
