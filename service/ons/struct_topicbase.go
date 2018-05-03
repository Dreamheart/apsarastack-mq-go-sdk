package ons

type TopicBase struct {
	AoneId			string	`json:"aoneId"`
	AoneUrl			string	`json:"aoneUrl"`
	Appkey			string	`json:"appkey"`
	Applier			string	`json:"applier"`
	ChannelId		int64	`json:"channelId"`
	ChannelName		string	`json:"channelName"`
	CreateTime		int64	`json:"createTime"`
	Id				int64	`json:"id"`
	Owner			string	`json:"owner"`
	PageSize		int		`json:"pageSize"`
	PageStart		int		`json:"pageStart"`
	Parent			string	`json:"parent"`
	RegionId		string	`json:"regionId"`
	RegionName		string	`json:"regionName"`
	Relation		int		`json:"relation"`
	RelationName	string	`json:"relationName"`
	Remark			string	`json:"remark"`
	RoleList		string	`json:"roleList"`
	RoleOwner		string	`json:"roleOwner"`
	Status			int		`json:"status"`
	StatusName		string	`json:"statusName"`
	Topic			string	`json:"topic"`
	UpdateTime		int64	`json:"updateTime"`
}
