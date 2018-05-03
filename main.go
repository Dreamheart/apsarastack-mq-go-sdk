package main

import (
	"github.com/Dreamheart/apsarastack-mq-go-sdk/service/ons"
)

func main() {
	ak := "accesskey"
	sk := "secretkey"
	endpoint := "mq.console.aliyun.com"
	client,err := ons.NewClient(ak,sk,endpoint)

	if err != nil {
		panic(err)
	}

	topic_name := "topic_create_by_gosdk"
	producer_id := "PID_test"
	consumer_id := "CID_test"
	regionId :="mq-private"

	desc_region_req := ons.CreateDescribeRegionsRequest()
	desc_region_resp,_ := client.DescribeRegions(desc_region_req)
	desc_region_resp.PrintSummary("DescribeRegion")
	desc_region_resp.PrintData()

	request := ons.CreateCreateTopicRequest(regionId, topic_name)
	create_topic_resp, _ := client.CreateTopic(request)
	create_topic_resp.PrintSummary("CreateTopic")


	desc_topic_req := ons.CreateDescribeTopicRequest(regionId, topic_name)
	desc_topic_resp, _ := client.DescribeTopic(desc_topic_req)
	desc_topic_resp.PrintSummary("DescribeTopic")
	desc_topic_resp.PrintData()


	pid_request := ons.CreateCreateProducerRequest(regionId, topic_name,producer_id)
	pid_response, _ := client.CreateProducer(pid_request)
	pid_response.PrintSummary("CreatePID")

	desc_pid_request := ons.CreateDescribeProducerRequest(regionId, topic_name, producer_id)
	desc_pid_resp,_ :=client.DescribeProducer(desc_pid_request)
	desc_pid_resp.PrintSummary("DescribeTopicProducer")
	desc_pid_resp.PrintData()


	delete_pid_req := ons.CreateDeleteProducerRequest(regionId, topic_name,producer_id)
	delete_pid_resp,_ :=client.DeleteProducer(delete_pid_req)
	delete_pid_resp.PrintSummary("DeletePID")


	cid_request := ons.CreateCreateConsumerRequest(regionId, topic_name,consumer_id)
	cid_response, _ := client.CreateConsumer(cid_request)
	cid_response.PrintSummary("CreateCID")

	desc_cid_req := ons.CreateDescribeConsumerRequest(regionId, topic_name, consumer_id)
	desc_cid_resp,_ := client.DescribeConsumer(desc_cid_req)
	desc_cid_resp.PrintSummary("DescribeConsumer")
	desc_cid_resp.PrintData()


	delete_cid_req := ons.CreateDeleteConsumerRequest(regionId, topic_name,consumer_id)
	delete_cid_resp,_ :=client.DeleteConsumer(delete_cid_req)
	delete_cid_resp.PrintSummary("DeleteCID")


	delete_topic_req := ons.CreateDeleteTopicRequest(regionId, topic_name)
	delete_topic_resp, _ := client.DeleteTopic(delete_topic_req)
	delete_topic_resp.PrintSummary("DeleteTopic")

	return
}




