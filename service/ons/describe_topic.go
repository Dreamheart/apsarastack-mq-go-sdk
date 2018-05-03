package ons

import (
	"mq_admin/mqsdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"net/http"
	"reflect"
	"fmt"
)

func (client *Client) DescribeTopic(request *DescribeTopicRequest) (response *DescribeTopicResponse, err error) {
	response = CreateDescribeTopicResponse()
	err = client.ProcessRequest(request,response)
	return
}

func CreateDescribeTopicRequest(regionId string, topic string) (request *DescribeTopicRequest) {
	request = &DescribeTopicRequest{
		CommonRequest: &mqsdk.CommonRequest{},
	}

	request.InitWithApiInfo("/json/topic/get",http.MethodGet)
	request.AddQueryParams("topic",topic)
	request.AddQueryParams("_regionId",regionId)

	return
}

func CreateDescribeTopicResponse() (response *DescribeTopicResponse) {
	response = &DescribeTopicResponse{
		CommonResponse: &CommonResponse{
			BaseResponse:&responses.BaseResponse{},
		},
	}
	return
}

type DescribeTopicRequest struct {
	*mqsdk.CommonRequest
	Topic string
}

type DescribeTopicResponse struct {
	*CommonResponse
	Data	[]Topic		`json:"data"`
}

func (response *DescribeTopicResponse) PrintData()  {
	entity_type := reflect.TypeOf(response.Data)
	for index, data := range response.Data{
		fmt.Printf("--- Print %s %d ---\n", entity_type.String(),index)
		t := reflect.TypeOf(data)
		v := reflect.ValueOf(data)
		count := v.NumField()
		for i := 0; i < count; i++ {
			f := v.Field(i)
			fmt.Printf("%20s : %v\n", t.Field(i).Name, f.Interface())
		}
	}
	fmt.Println("")
}



