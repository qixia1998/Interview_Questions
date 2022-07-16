package main

import (
	"Interview_Questions/GoWeb/rabbitmq/RabbitMQ"
	"fmt"
	"strconv"
	"time"
)

func main() {
	qixiaOne := RabbitMQ.NewRabbitMQRouting("exQixia", "qixia_one")
	qixiaTwo := RabbitMQ.NewRabbitMQRouting("exQixia", "qixia_two")
	for i := 0; i <= 10; i++ {
		qixiaOne.PublishRouting("Hello qixia one!" + strconv.Itoa(i))
		qixiaTwo.PublishRouting("Hello qixia two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
