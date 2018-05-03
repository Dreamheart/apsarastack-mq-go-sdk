package ons

import (
	"github.com/Dreamheart/apsarastack-mq-go-sdk/mqsdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"net/http"
)

func (client *Client) CreateConsumer(request *CreateConsumerRequest) (response *CreateConsumerResponse, err error) {
	response = CreateCreateConsumerResponse()
	err = client.ProcessRequest(request,response)
	return
}

func CreateCreateConsumerRequest(regionId string, topic string,ConsumerId string) (request *CreateConsumerRequest) {
	request = &CreateConsumerRequest{
		CommonRequest: &mqsdk.CommonRequest{},
	}

	request.InitWithApiInfo("/json/subscription/create",http.MethodPost)
	request.AddQueryParams("topic",topic)
	request.AddQueryParams("ConsumerId",ConsumerId)
	request.AddQueryParams("_regionId",regionId)

	return
}

func CreateCreateConsumerResponse() (response *CreateConsumerResponse) {
	response = &CreateConsumerResponse{
		CommonResponse: &CommonResponse{
			BaseResponse:&responses.BaseResponse{},
		},
	}
	return
}

type CreateConsumerRequest struct {
	*mqsdk.CommonRequest
	Topic string
}

type CreateConsumerResponse struct {
	*CommonResponse
}

