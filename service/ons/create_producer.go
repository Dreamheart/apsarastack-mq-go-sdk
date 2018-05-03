package ons

import (
	"github.com/Dreamheart/apsarastack-mq-go-sdk/mqsdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"net/http"
)

func (client *Client) CreateProducer(request *CreateProducerRequest) (response *CreateProducerResponse, err error) {
	response = CreateCreateProducerResponse()
	err = client.ProcessRequest(request,response)
	return
}

func CreateCreateProducerRequest(regionId string, topic string,producerId string) (request *CreateProducerRequest) {
	request = &CreateProducerRequest{
		CommonRequest: &mqsdk.CommonRequest{},
	}

	request.InitWithApiInfo("/json/publish/create",http.MethodPost)
	request.AddQueryParams("topic",topic)
	request.AddQueryParams("ProducerId",producerId)
	request.AddQueryParams("_regionId",regionId)

	return
}

func CreateCreateProducerResponse() (response *CreateProducerResponse) {
	response = &CreateProducerResponse{
		CommonResponse: &CommonResponse{
			BaseResponse:&responses.BaseResponse{},
		},
	}
	return
}

type CreateProducerRequest struct {
	*mqsdk.CommonRequest
	Topic string
}

type CreateProducerResponse struct {
	*CommonResponse
}

