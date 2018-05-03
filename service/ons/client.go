package ons

import (
	"github.com/Dreamheart/apsarastack-mq-go-sdk/mqsdk"
)

type Client struct {
	mqsdk.Client
}

func NewClient(ak string, sk string, endpoint string) (client *Client, err error) {
	client = &Client{}
	client.InitWithAccessKey(ak,sk,endpoint)
	return
}
