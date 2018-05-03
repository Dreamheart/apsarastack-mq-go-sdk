package ons

import (
	"github.com/Dreamheart/apsarastack-mq-go-sdk/mqsdk"
	"net/http"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"reflect"
	"fmt"
)

func (client *Client) DescribeProducer(request *DescribeProducerRequest) (response *DescribeProducerResponse, err error) {
	response = CreateDescribeProducerResponse()
	err = client.ProcessRequest(request,response)
	return
}

func CreateDescribeProducerRequest(regionId string, topic string, producerId string) (request *DescribeProducerRequest) {
	request = &DescribeProducerRequest{
		CommonRequest: &mqsdk.CommonRequest{},
	}

	request.InitWithApiInfo("/json/publish/get",http.MethodGet)
	request.AddQueryParams("topic",topic)
	request.AddQueryParams("_regionId",regionId)
	request.AddQueryParams("topic",topic)
	request.AddQueryParams("producerId",producerId)

	return
}

func CreateDescribeProducerResponse() (response *DescribeProducerResponse) {
	response = &DescribeProducerResponse{
		CommonResponse: &CommonResponse{
			BaseResponse:&responses.BaseResponse{},
		},
	}
	return
}

type DescribeProducerRequest struct {
	*mqsdk.CommonRequest
	Topic string
}

type DescribeProducerResponse struct {
	*CommonResponse
	Data	[]TopicProducer		`json:"data"`
}

func (response *DescribeProducerResponse) PrintData()  {
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
