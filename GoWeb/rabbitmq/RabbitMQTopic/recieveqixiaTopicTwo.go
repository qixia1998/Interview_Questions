package main

import "Interview_Questions/GoWeb/rabbitmq/RabbitMQ"

func main() {
	qixiaOne := RabbitMQ.NewRabbitMQTopic("exQixaTopic", "qixia.*.two")
	qixiaOne.RecieveTopic()
}
