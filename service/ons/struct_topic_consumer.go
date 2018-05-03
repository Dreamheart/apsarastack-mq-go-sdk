package ons

type TopicConsumer struct {
	*TopicBase
	ConsumerId			string	`json:"consumerId"`
}
