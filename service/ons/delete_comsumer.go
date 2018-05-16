package ons

import (
	"github.com/Dreamheart/apsarastack-mq-go-sdk/mqsdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"net/http"
)

func (client *Client) DeleteConsumer(request *DeleteConsumerRequest) (response *DeleteConsumerResponse, err error) {
	response = CreateDeleteConsumerResponse()
	err = client.ProcessRequest(request,response)
	return
}

func CreateDeleteConsumerRequest(regionId string, topic string,ConsumerId string) (request *DeleteConsumerRequest) {
	request = &DeleteConsumerRequest{
		CommonRequest: &mqsdk.CommonRequest{},
	}

	request.InitWithApiInfo("/json/subscription/delete",http.MethodPost)
	request.AddQueryParams("topic",topic)
	request.AddQueryParams("consumerId",ConsumerId)
	request.AddQueryParams("_regionId",regionId)

	return
}

func CreateDeleteConsumerResponse() (response *DeleteConsumerResponse) {
	response = &DeleteConsumerResponse{
		CommonResponse: &CommonResponse{
			BaseResponse:&responses.BaseResponse{},
		},
	}
	return
}

type DeleteConsumerRequest struct {
	*mqsdk.CommonRequest
	Topic string
}

type DeleteConsumerResponse struct {
	*CommonResponse
}

