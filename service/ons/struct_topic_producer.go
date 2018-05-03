package ons

type TopicProducer struct {
	*TopicBase
	ProducerId			string	`json:"producerId"`
}
