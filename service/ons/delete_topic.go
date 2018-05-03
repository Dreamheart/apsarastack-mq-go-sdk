package ons

import (
	"mq_admin/mqsdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"net/http"
)

func (client *Client) DeleteTopic(request *DeleteTopicRequest) (response *DeleteTopicResponse, err error) {
	response = CreateDeleteTopicResponse()
	err = client.ProcessRequest(request,response)
	return
}

func CreateDeleteTopicRequest(regionId string, topic string) (request *DeleteTopicRequest) {
	request = &DeleteTopicRequest{
		CommonRequest: &mqsdk.CommonRequest{},
	}

	request.InitWithApiInfo("/json/topic/delete",http.MethodPost)
	request.AddQueryParams("topic",topic)
	request.AddQueryParams("_regionId",regionId)

	return
}

func CreateDeleteTopicResponse() (response *DeleteTopicResponse) {
	response = &DeleteTopicResponse{
		CommonResponse: &CommonResponse{
			BaseResponse:&responses.BaseResponse{},
		},
	}
	return
}

type DeleteTopicRequest struct {
	*mqsdk.CommonRequest
	Topic string
}

type DeleteTopicResponse struct {
	*CommonResponse
}

