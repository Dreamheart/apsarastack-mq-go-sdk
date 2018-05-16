package mqsdk

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/utils"
	"strings"
	"io"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

type ICommonRequest interface {
	AddQueryParams(key string, value string)
	AddFormParams(key string, value string)
	AddHeaders(key string, value string)
	GetQueryParams() map[string]string
	GetFormParams() map[string]string
	GetAcceptFormat() string
	BuildQueries() string
	BuildUrl() string
	InitWithApiInfo(path string, method string)
	GetBodyReader() io.Reader
	GetHeaders() map[string]string
	GetMethod() string
	GetDomain() string
	SetDomain(domain string)
}


type CommonRequest struct{
	Scheme		string
	Domain		string
	Method		string
	AcceptFormat	string
	PathPattern	string
	Headers     map[string]string
	QueryParams map[string]string
	FormParams  map[string]string

	queries		string
}

func NewCommonRequest(path string, method string) (request *CommonRequest){
	request = &CommonRequest{}
	request.Init()
	request.PathPattern = path
	request.Method = method
	
	return 
}

func (request *CommonRequest)Init(){
	request.Scheme = requests.HTTP
	request.Domain = ""
	request.AcceptFormat = "JSON"

	request.Headers = make(map[string]string)
	request.QueryParams = make(map[string]string)
	request.FormParams = make(map[string]string)
}

func (request *CommonRequest) GetHeaders() map[string]string {
	return request.Headers
}

func (request *CommonRequest) GetMethod() string {
	return request.Method
}

func (request *CommonRequest) GetDomain() string {
	return request.Domain
}

func (request *CommonRequest) SetDomain(domain string){
	request.Domain = domain
}

func (request *CommonRequest) AddQueryParams(key string, value string)  {
	request.QueryParams[key] = value
}

func (request *CommonRequest) AddFormParams(key string, value string)  {
	request.FormParams[key] = value
}

func (request *CommonRequest) AddHeaders(key string, value string)  {
	request.Headers[key] = value
}

func (request *CommonRequest) GetQueryParams() map[string]string {
	return request.QueryParams
}

func (request *CommonRequest) GetFormParams() map[string]string {
	return request.FormParams
}

func (request *CommonRequest) GetAcceptFormat() string {
	return request.AcceptFormat
}

func (request *CommonRequest) BuildQueries() string {
	request.queries = "/?" + utils.GetUrlFormedMap(request.QueryParams)
	return request.queries
}


func (request *CommonRequest) BuildUrl() string {
	//return strings.Replace(request.Domain + request.PathPattern+ request.BuildQueries(), "&", "\\&",-1)
	return request.Domain + request.PathPattern+ request.BuildQueries()
}

func (request *CommonRequest) InitWithApiInfo(path string, method string) {
	request.Init()
	request.PathPattern = path
	request.Method = method
}


func (request *CommonRequest) GetBodyReader() io.Reader {
	if request.FormParams != nil && len(request.FormParams) > 0 {
		formString := utils.GetUrlFormedMap(request.FormParams)
		return strings.NewReader(formString)
	} else {
		return strings.NewReader("")
	}
}
