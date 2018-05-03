package ons

type Topic struct {
	*TopicBase
	OrderType		int		`json:"orderType"`
}