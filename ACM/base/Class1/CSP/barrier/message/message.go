package message

import "fmt"

type Consumer struct{}

func (c *Consumer) Exit() {
	fmt.Println("Consumer Exit")
}

func (c *Consumer) Join() {
	fmt.Println("Consumer Join")
}

type Producer struct{}

func (p *Producer) Exit() {
	fmt.Println("Producer Exit")
}

func (p *Producer) Join() {
	fmt.Println("Producer Join")
}

var consumerInstance *Consumer
var producerInstance *Producer

func init() {
	consumerInstance = &Consumer{}
	producerInstance = &Producer{}
}

func ConsumerInst() *Consumer {
	return consumerInstance
}

func ProducerInst() *Producer {
	return producerInstance
}
