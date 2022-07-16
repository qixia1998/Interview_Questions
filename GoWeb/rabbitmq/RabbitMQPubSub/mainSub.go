package main

import "Interview_Questions/GoWeb/rabbitmq/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("" + "newProduct")
	rabbitmq.RecieveSub()
}
