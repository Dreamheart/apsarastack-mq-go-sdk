package ons

import (
	"mq_admin/mqsdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"net/http"
)

func (client *Client) CreateTopic(request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
	response = CreateCreateTopicResponse()
	err = client.ProcessRequest(request,response)
	return
}

func CreateCreateTopicRequest(regionId string, topic string) (request *CreateTopicRequest) {
	request = &CreateTopicRequest{
		CommonRequest: &mqsdk.CommonRequest{},
	}

	request.InitWithApiInfo("/json/topic/create",http.MethodPost)
	request.AddQueryParams("topic",topic)
	request.AddQueryParams("_regionId",regionId)


	return
}

func CreateCreateTopicResponse() (response *CreateTopicResponse) {
	response = &CreateTopicResponse{
		CommonResponse: &CommonResponse{
			BaseResponse:&responses.BaseResponse{},
		},
	}
	return
}

type CreateTopicRequest struct {
	*mqsdk.CommonRequest
	Topic string
}

type CreateTopicResponse struct {
	*CommonResponse
}

