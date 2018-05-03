package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/Dreamheart/apsarastack-mq-go-sdk/mqsdk"
	"net/http"
	"reflect"
	"fmt"
)

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
	response = CreateDescribeRegionsResponse()
	err = client.ProcessRequest(request,response)
	return
}

// DescribeRegionsRequest is the request struct for api DescribeRegions
type DescribeRegionsRequest struct {
	*mqsdk.CommonRequest
}

func CreateDescribeRegionsRequest() (request *DescribeRegionsRequest) {
	request = &DescribeRegionsRequest{
		CommonRequest: &mqsdk.CommonRequest{},
	}

	request.InitWithApiInfo("/json/region/list",http.MethodGet)

	return
}

// DescribeRegionsResponse is the response struct for api DescribeRegions
type DescribeRegionsResponse struct {
	*CommonResponse
	Data   []Region `json:"data"`
}

func CreateDescribeRegionsResponse() (response *DescribeRegionsResponse) {
	response = &DescribeRegionsResponse{
		CommonResponse: &CommonResponse{
			BaseResponse:&responses.BaseResponse{},
		},
	}
	return
}

func (response *DescribeRegionsResponse) PrintData()  {
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

