package messageQueue

import (
	"XChatRoom/data/dataImpl"
	"XChatRoom/model"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"strconv"
	"strings"
)


func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s", msg, err)
	}
}

/*
	Desc:实例化一个RabbitMq对象
 */
func CreateConnectionRabbitMQ() {

	ConnAMQP, err := amqp.Dial("amqp://lyt:root@localhost:5672/")
	failOnError(err, "Fail to connect to RabbitMQ.")
	fmt.Println("ConnAMQP == nil ?", ConnAMQP == nil)
	defer ConnAMQP.Close()
}

/*
	Desc: 发送一条消息
	Input:
		from:发件人
		target:收件人
		content:消息内容
 */
func PublishMessage(from, target, content string) {

	conn, err := amqp.Dial("amqp://lyt:root@localhost:5672/")
	failOnError(err, "Fail to connect to RabbitMQ.")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Fail to connect to Channel.")
	defer ch.Close()

	//声明Exchange
	err = ch.ExchangeDeclare(
		"testcontent",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Fail to declare exchange.")


	err = ch.Publish(
		"testcontent", //聊天消息接收队列的exchange
		target,				//接受消息者的用户名作为route-key
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(from + ":" + content),//发布格式——发件人:消息内容
		})
	failOnError(err, "Fail to publish queue [content].")

	log.Printf("<amqp>: from[%s] sent to target[%s] content[%s]",from, target, content)

}

func BgRecMessage(userId uint64, username string) {

	conn, err := amqp.Dial("amqp://lyt:root@localhost:5672/")
	failOnError(err, "Fail to connect to RabbitMQ.")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Fail to connect to Channel.")
	defer ch.Close()

	//声明Exchange
	err = ch.ExchangeDeclare(
		"testcontent",
		"direct",
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
		username,//route key = username
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


	//forever := make(chan bool)
	go func() {
		for d := range msgs {

			bodySplitStr := strings.Split(string(d.Body[:]), ":")	//拆解消息，通过":"分隔符
			fromId, _ := strconv.ParseUint(bodySplitStr[0], 10, 64)
			log.Printf("[AMQP] %s", d.Body)
			mes := dataImpl.Message{
				Id:      userId,
				Type:    "normal",
				Content: bodySplitStr[1],	//[1]:content
				Source:  fromId,	//[0]:发件人userId
				Target:  userId,
			}

			model.AddMessage(model.Client, mes)//添加到Redis UserName: Message<K,V>

		}
	}()
	//<-forever
}