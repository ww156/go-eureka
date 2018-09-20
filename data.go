package eureka

import "strconv"

type App struct {
	Application *Instance `json:"application" xml:"application"`
}

type Instance struct {
	InstanceId           string                 `xml:"instanceId" json:"instanceId"`
	HostName             string                 `xml:"hostName" json:"hostName"`
	App                  string                 `xml:"app" json:"app"`
	IPAddr               string                 `xml:"ipAddr" json:"ipAddr"`
	VipAddress           string                 `xml:"vipAddress" json:"vipAddress"`
	SecureVipAddress     string                 `xml:"secureVipAddress" json:"secureVipAddress"`
	Status               string                 `xml:"status" json:"status"`
	Port                 *Port                  `xml:"port" json:"port"`
	SecurePort           *Port                  `xml:"securePort" json:"securePort"`
	HomePageUrl          string                 `xml:"homePageUrl" json:"homePageUrl"`
	StatusPageUrl        string                 `xml:"statusPageUrl" json:"statusPageUrl"`
	HealthCheckUrl       string                 `xml:"healthCheckUrl" json:"healthCheckUrl"`
	SecureHealthCheckUrl string                 `xml:"secure_health_check_url" json:"secure_health_check_url"`
	DataCenterInfo       *DataCenterInfo        `xml:"dataCenterInfo" json:"dataCenterInfo"`
	LeaseInfo            *LeaseInfo             `xml:"lease_info" json:"lease_info"`
	Metadata             map[string]interface{} `xml:"metadata" json:"metadata"`
}
type Port struct {
	Port   int  `json:"$"`
	Enable bool `json:"@enabled"`
}

type DataCenterInfo struct {
	Name  string `json:"name" xml:"name"`
	Class string `json:"@class"`
}

type LeaseInfo struct {
	RenewalIntervalInSecs int   `xml:"renewal_interval_in_secs" json:"renewal_interval_in_secs"`
	DurationInSecs        int   `xml:"duration_in_secs" json:"duration_in_secs"`
	RegistrationTimestamp int64 `xml:"registration_timestamp" json:"registration_timestamp"`
	LastRenewalTimestamp  int64 `xml:"last_renewal_timestamp" json:"last_renewal_timestamp"`
	RenewalTimestamp      int64 `xml:"renewal_timestamp" json:"renewal_timestamp"`
	EvictionTimestamp     int64 `xml:"eviction_timestamp" json:"eviction_timestamp"`
	ServiceUpTimestamp    int64 `xml:"service_up_timestamp" json:"service_up_timestamp"`
}

func (i *Instance) Id() string {
	if i.InstanceId != "" {
		return i.InstanceId
	}
	return i.HostName + ":" + i.App + ":" + strconv.Itoa(i.Port.Port)
}
