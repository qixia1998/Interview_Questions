package main

import (
	"Interview_Questions/GoWeb/rabbitmq/RabbitMQ"
	"fmt"
	"strconv"
	"time"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "qixiaSimple")

	for i := 0; i <= 100; i++ {
		rabbitmq.PublishSimple("Hello qixia!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
