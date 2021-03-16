package ssov3

import (
	"iotdemo/module/httputil"
	"iotdemo/pkg/utils/syslog"
	"errors"
	"github.com/json-iterator/go"
	"os"
)

var SSoUrl string = os.Getenv("sso_url")
func init(){
	LOG.Info("ssourl:", SSoUrl)
}

type UserRole struct {
	RoleName         string `json:"roleName"`
	Workspace        string `json:"workspace"`
	Namespace        string `json:"namespace"`
	Cluster          string `json:"cluster"`
	Datacenter       string `json:"datacenter"`
	SubscriptionName string `json:"SubscriptionName"`
	SubscriptionId   string `json:"SubscriptionId"`
}

type UserSubscriptionDto struct {
	SubscriptionId    string `json:"subscriptionId"`
	SubscriptionName       string `json:"subscriptionName"  xorm:"name"`
	Subscription_role string `json:"subscriptionRole"`
	Company    string `json:"company"`
}


type SignedInUser struct {
	Id                          string          `json:"id"`
	Address                     string          `json:"address"`
	Alternative_email           string          `json:"alternativeEmail"`
	Avatar                      string          `json:"avatar"`
	City                        string          `json:"city"`
	Company                     string          `json:"company"`
	Contact_phone               string          `json:"contactPhone"`
	Country                     string          `json:"country"`
	Creation_time               int64           `json:"creationTime"`
	Creator                     string          `json:"creator"`
	Expiration_time             int64           `json:"expirationTime"`
	First_name                  string          `json:"firstName"`
	Groups                      []string     `json:"groups"`
	Industry                    string          `json:"industry"`
	Last_city                   string          `json:"lastCity"`
	Last_ip                     string          `json:"lastIp"`
	Last_lat                    float64         `json:"lastLat"`
	Last_long                   float64         `json:"lastLong"`
	Last_modified_pwd_time      int64           `json:"lastModifiedPwdTime"`
	Last_modified_time          int64           `json:"lastModifiedTime"`
	Last_name                   string          `json:"lastName"`
	Last_signed_in_time         int64           `json:"lastSignedInTime"`
	Mobile_phone                string          `json:"mobilePhone"`
	Office_phone                string          `json:"officePhone"`
	Origin                      string          `json:"origin"`
	Postal_code                 string          `json:"postalCode"`
	Role                        string          `json:"role"`
	Signed_in_frequency_counter int64           `json:"signedInFrequencyCounter"`
	Username                    string          `json:"username"`
	Total_signed_in_times       int64           `json:"totalSignedInTimes"`
	Salt                        string          `json:"salt"`
	Roles                       []*UserRole     `json:"roles"`
	UserSubscriptions           []*UserSubscriptionDto `json:"userSubscriptions"`
	Scopes                      []string        `json:"scopes"`
	Status                      string          `json:"status"`
}


func GetUsersMe(token string) (*SignedInUser, error) {
	var url = SSoUrl+"users/me"
	var header map[string]string
	header = make(map[string]string)
	authString := token
	header["Content-type"] = "application/json"
	header["Authorization"] = authString
	response, status := httputil.HttpGet(url, header, "")
	if status != 200 {
		error := errors.New(response)
		return nil, error
	}
	var userInfo = new(SignedInUser)
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	bResp := []byte(response)
	ret := json_iterator.Unmarshal(bResp, userInfo)
	//fmt.Printf("userInfo:%+v\n", userInfo)
	if ret != nil {
		error := errors.New("json unmarshal fail")
		return nil, error
	} else {
		return userInfo, nil
	}
}
