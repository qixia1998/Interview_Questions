package main

import "Interview_Questions/GoWeb/rabbitmq/RabbitMQ"

func main() {
	qixiaOne := RabbitMQ.NewRabbitMQRouting("exQixia", "qixia_one")
	qixiaOne.RecieveRouting()
}
