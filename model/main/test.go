package main

import (
	"XChatRoom/messageQueue"
)

func main() {
	messageQueue.PublishMessage("lyt", "xiaoming", "helloThisisMyTest")
}

