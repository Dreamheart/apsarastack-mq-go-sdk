package mqsdk

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/signers"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth"
	"sort"
	"bytes"
	"net/http"
	"time"
	"strconv"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type Client struct {
	AccessKey  	string
	SecretKey  	string
	Endpoint   	string

	signer		*signers.AccessKeySigner
	httpClient	*http.Client
}

func NewClient(ak string, sk string, endpoint string) (client *Client, err error) {
	client = &Client{}
	err = client.InitWithAccessKey(ak,sk,endpoint)
	return
}

func (client *Client) InitWithAccessKey(ak string, sk string, endpoint string) (err error) {
	client.AccessKey = ak
	client.SecretKey = sk
	client.Endpoint = endpoint

	credential := &credentials.AccessKeyCredential{
		AccessKeyId:     ak,
		AccessKeySecret: sk,
	}
	client.signer, err = signers.NewAccessKeySigner(credential)

	client.httpClient = &http.Client{}
	return
}

func (client *Client) ProcessRequest(request ICommonRequest,response responses.AcsResponse) (err error) {

	request.SetDomain(client.Endpoint)
	httpRequest, err := buildHttpRequest(request, client.signer)
	if err != nil{
		panic(err)
	}

	var httpResponse *http.Response
	httpResponse, err = client.httpClient.Do(httpRequest)

	if err != nil {
		panic(err)
	}

	err = responses.Unmarshal(response, httpResponse, request.GetAcceptFormat())
	//// wrap server errors
	//if serverErr, ok := err.(*errors.ServerError); ok {
	//	var wrapInfo = map[string]string{}
	//	wrapInfo["StringToSign"] = request.GetStringToSign()
	//	err = errors.WrapServerError(serverErr, wrapInfo)
	//}

	//response = NewCommonResponse("Not Implement ...")
	return
}

func buildHttpRequest(request ICommonRequest, signer auth.Signer) (httpRequest *http.Request, err error) {
	completeRpcSignParams(request, signer)
	//构造签名串
	stringToSign := buildRpcStringToSign(request, signer)
	//签名
	signature := signer.Sign(stringToSign,"")

	request.AddQueryParams("_signature",signature)

	requestMethod := request.GetMethod()
	requestUrl := request.BuildUrl()
	//fmt.Println(requestUrl)
	body := request.GetBodyReader()
	httpRequest, err = http.NewRequest(requestMethod, requestUrl, body)
	if err != nil {
		return
	}
	for key, value := range request.GetHeaders() {
		httpRequest.Header[key] = []string{value}
	}
	return
}
func completeRpcSignParams(request ICommonRequest, signer auth.Signer) (err error) {
	queryParams := request.GetQueryParams()
	queryParams["_accesskey"], err = signer.GetAccessKeyId()
	queryParams["_platform"] = "onsConsole"
	queryParams["__preventCache"] = strconv.FormatInt(time.Now().Unix() * 1000,10)

	if err != nil {
		return
	}

	//request.GetHeaders()["Content-Type"] = requests.Form
	//formString := utils.GetUrlFormedMap(request.GetFormParams())
	//request.SetContent([]byte(formString))

	return
}

func buildRpcStringToSign(request ICommonRequest, signer auth.Signer) (stringToSign string) {

	signParams := make(map[string]string)
	for key, value := range request.GetQueryParams() {
		signParams[key] = value
	}
	for key, value := range request.GetFormParams() {
		signParams[key] = value
	}

	// sort params by key
	var paramKeySlice []string
	for key := range signParams {
		paramKeySlice = append(paramKeySlice, key)
	}
	sort.Strings(paramKeySlice)
	var buf bytes.Buffer
	for _,key := range paramKeySlice {
		buf.WriteString(signParams[key])
	}
	stringToSign = buf.String()
	return
}
