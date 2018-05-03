package ons

import (
	"github.com/Dreamheart/apsarastack-mq-go-sdk/mqsdk"
	"net/http"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"reflect"
	"fmt"
)

func (client *Client) DescribeConsumer(request *DescribeConsumerRequest) (response *DescribeConsumerResponse, err error) {
	response = CreateDescribeConsumerResponse()
	err = client.ProcessRequest(request,response)
	return
}

func CreateDescribeConsumerRequest(regionId string, topic string, consumerId string) (request *DescribeConsumerRequest) {
	request = &DescribeConsumerRequest{
		CommonRequest: &mqsdk.CommonRequest{},
	}

	request.InitWithApiInfo("/json/subscription/get",http.MethodGet)
	request.AddQueryParams("topic",topic)
	request.AddQueryParams("_regionId",regionId)
	request.AddQueryParams("topic",topic)
	request.AddQueryParams("consumerId",consumerId)

	return
}

func CreateDescribeConsumerResponse() (response *DescribeConsumerResponse) {
	response = &DescribeConsumerResponse{
		CommonResponse: &CommonResponse{
			BaseResponse:&responses.BaseResponse{},
		},
	}
	return
}

type DescribeConsumerRequest struct {
	*mqsdk.CommonRequest
	Topic string
}

type DescribeConsumerResponse struct {
	*CommonResponse
	Data	[]TopicConsumer		`json:"data"`
}

func (response *DescribeConsumerResponse) PrintData()  {
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
