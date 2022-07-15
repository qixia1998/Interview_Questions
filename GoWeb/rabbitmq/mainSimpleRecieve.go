package main

import "Interview_Questions/GoWeb/rabbitmq/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("qixiaSimple")
	rabbitmq.ConsumeSimple()
}
