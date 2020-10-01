package main

import (
	"container/list"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s", msg, err)
	}
}


func main() {
	ConsumeMessage("xiaoming")
}

func ConsumeMessage(username string) *list.List{

	conn, err := amqp.Dial("amqp://lyt:root@localhost:5672/")
	failOnError(err, "Fail to connect to RabbitMQ.")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Fail to connect to Channel.")
	defer ch.Close()

	//声明Exchange
	err = ch.ExchangeDeclare(
		"testcontent",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare an exchange")


	queue, err := ch.QueueDeclare(
		"default",
		false,
		false,
		true,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		queue.Name,//队列名
		//username,//route key = username
		"",
		"testcontent",//exhange = content(聊天消息内容exchange)
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")
	log.Printf("Binding queue %s to exchange %s with routing key %s",
		queue.Name, "content", username)


	//注册消费——所有指向该收件人的消息
	//msgs, err := ch.Consume(
	msgs, err := ch.Consume(
		queue.Name,
		"defaultConsumer",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

forever := make(chan bool)
	res := list.New()
	fmt.Println("len(msgs) =", len(msgs))

	go func() {
		for d := range msgs {
			log.Printf("[AMQP] %s", d.Body)
		}
	}()

<-forever


	//mock
	//res.PushBack([]byte{'x','i','a','o','m','i','n','g',':','h','e','e','l','o'})
	return res
}

