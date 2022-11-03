package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topic", //name of the exchange
		"topic",      //type
		true,         //durable?
		false,        //auto-deleted?
		false,        //internal?
		false,        //no wait?
		nil,          //arguments?
	)
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    //Name
		false, //durable?
		false, // delete it when unsused?
		true,  //exclusive?
		false, // is it no wait?
		nil,   //arguments?
	)
}
