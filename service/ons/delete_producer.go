package ons

import (
	"github.com/Dreamheart/apsarastack-mq-go-sdk/mqsdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"net/http"
)

func (client *Client) DeleteProducer(request *DeleteProducerRequest) (response *DeleteProducerResponse, err error) {
	response = CreateDeleteProducerResponse()
	err = client.ProcessRequest(request,response)
	return
}

func CreateDeleteProducerRequest(regionId string, topic string,producerId string) (request *DeleteProducerRequest) {
	request = &DeleteProducerRequest{
		CommonRequest: &mqsdk.CommonRequest{},
	}

	request.InitWithApiInfo("/json/publish/delete",http.MethodPost)
	request.AddQueryParams("topic",topic)
	request.AddQueryParams("ProducerId",producerId)
	request.AddQueryParams("_regionId",regionId)

	return
}

func CreateDeleteProducerResponse() (response *DeleteProducerResponse) {
	response = &DeleteProducerResponse{
		CommonResponse: &CommonResponse{
			BaseResponse:&responses.BaseResponse{},
		},
	}
	return
}

type DeleteProducerRequest struct {
	*mqsdk.CommonRequest
	Topic string
}

type DeleteProducerResponse struct {
	*CommonResponse
}

