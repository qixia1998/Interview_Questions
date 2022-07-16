package main

import (
	"Interview_Questions/GoWeb/rabbitmq/RabbitMQ"
	"fmt"
	"strconv"
	"time"
)

func main() {
	qixiaOne := RabbitMQ.NewRabbitMQTopic("exQixiaTopic", "qixia.topic.one")
	qixiaTwo := RabbitMQ.NewRabbitMQTopic("exQixiaTopic", "qixia.topic.two")
	for i := 0; i <= 10; i++ {
		qixiaOne.PublishTopic("Hello qixia topic One!" + strconv.Itoa(i))
		qixiaTwo.PublishTopic("Hello qixia topic Two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
