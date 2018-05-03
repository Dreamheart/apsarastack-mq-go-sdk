package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"fmt"
)

type CommonResponse struct {
	*responses.BaseResponse
	Code 			int		`json:"code"`
	CurrentPage		int 	`json:"currentPage"`
	DynamicCode		string 	`json:"dynamicCode"`
	DynamicMessage	string 	`json:"dynamicMessage"`
	Message			string 	`json:"message"`
	PageSize		int 	`json:"pageSize"`
	RequestId		string 	`json:"requestId"`
	Status			string 	`json:"status"`
	Success			bool 	`json:"success"`
	Total			int 	`json:"total"`
}

func (response *CommonResponse) PrintSummary(tag string)  {
	fmt.Printf("-------------%s-------------\n",tag)
	fmt.Println("Success:",response.Success)
	fmt.Println("Status:",response.Status)
	fmt.Println("Message:",response.Message)

	fmt.Println("")
}

func (response *CommonResponse) PrintData()  {
	fmt.Println("Not Implement ...")
}