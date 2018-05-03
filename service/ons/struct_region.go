package ons

type Region struct {
	RegionId  string `json:"regionId" xml:"RegionId"`
	RegionName string `json:"regionName" xml:"LocalName"`
	Status    string `json:"status" xml:"Status"`
}