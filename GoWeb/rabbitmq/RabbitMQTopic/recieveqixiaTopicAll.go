package main

import "Interview_Questions/GoWeb/rabbitmq/RabbitMQ"

func main() {
	qixiaOne := RabbitMQ.NewRabbitMQTopic("exQixiaTopic", "#")
	qixiaOne.RecieveTopic()
}
