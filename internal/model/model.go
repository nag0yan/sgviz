package model

type CLIResponse struct {
	SecurityGroups []SecurityGroup `json:"SecurityGroups"`
	NextToken      string          `json:"NextToken"`
}
type IPRange struct {
	CidrIP      string `json:"CidrIp"`
	Description string `json:"Description"`
}
type Ipv6Range struct {
	CidrIpv6    string `json:"CidrIpv6"`
	Description string `json:"Description"`
}
type PrefixListId struct {
	Description  string `json:"Description"`
	PrefixListID string `json:"PrefixListId"`
}
type UserIDGroupPair struct {
	Description            string `json:"Description"`
	GroupID                string `json:"GroupId"`
	GroupName              string `json:"GroupName"`
	PeeringStatus          string `json:"PeeringStatus"`
	UserID                 string `json:"UserId"`
	VpcID                  string `json:"VpcId"`
	VpcPeeringConnectionID string `json:"VpcPeeringConnectionId"`
}
type IPPermission struct {
	FromPort         int               `json:"FromPort"`
	IPProtocol       string            `json:"IpProtocol"`
	IPRanges         []IPRange         `json:"IpRanges"`
	Ipv6Ranges       []Ipv6Range       `json:"Ipv6Ranges"`
	PrefixListIds    []PrefixListId    `json:"PrefixListIds"`
	ToPort           int               `json:"ToPort"`
	UserIDGroupPairs []UserIDGroupPair `json:"UserIdGroupPairs"`
}
type IPPermissionEgress struct {
	FromPort         int               `json:"FromPort"`
	IPProtocol       string            `json:"IpProtocol"`
	IPRanges         []IPRange         `json:"IpRanges"`
	Ipv6Ranges       []Ipv6Range       `json:"Ipv6Ranges"`
	PrefixListIds    []PrefixListId    `json:"PrefixListIds"`
	ToPort           int               `json:"ToPort"`
	UserIDGroupPairs []UserIDGroupPair `json:"UserIdGroupPairs"`
}
type Tag struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}
type SecurityGroup struct {
	Description         string               `json:"Description"`
	GroupName           string               `json:"GroupName"`
	IPPermissions       []IPPermission       `json:"IpPermissions"`
	OwnerID             string               `json:"OwnerId"`
	GroupID             string               `json:"GroupId"`
	IPPermissionsEgress []IPPermissionEgress `json:"IpPermissionsEgress"`
	Tags                []Tag                `json:"Tags"`
	VpcID               string               `json:"VpcId"`
}
