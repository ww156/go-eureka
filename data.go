package eureka

import "strconv"

type Application struct {
	Application ResApp `xml:"application" json:"application"`
}

type ResApp struct {
	Name     string     `xml:"name" json:"name"`
	Instance []Instance `xml:"instance" json:"instance"`
}

// 注册App
type App struct {
	Instance *Instance `xml:"instance" json:"instance"`
}

type Instance struct {
	InstanceId                    string                 `xml:"instanceId" json:"instanceId"`
	HostName                      string                 `xml:"hostName" json:"hostName"`
	App                           string                 `xml:"app" json:"app"`
	IPAddr                        string                 `xml:"ipAddr" json:"ipAddr"`
	VipAddress                    string                 `xml:"vipAddress" json:"vipAddress"`
	SecureVipAddress              string                 `xml:"secureVipAddress" json:"secureVipAddress"`
	Status                        string                 `xml:"status" json:"status"`
	Port                          *Port                  `xml:"port" json:"port"`
	SecurePort                    *Port                  `xml:"securePort" json:"securePort"`
	HomePageUrl                   string                 `xml:"homePageUrl" json:"homePageUrl"`
	StatusPageUrl                 string                 `xml:"statusPageUrl" json:"statusPageUrl"`
	HealthCheckUrl                string                 `xml:"healthCheckUrl" json:"healthCheckUrl"`
	SecureHealthCheckUrl          string                 `xml:"secureHealthCheckUrl" json:"secureHealthCheckUrl"`
	DataCenterInfo                *DataCenterInfo        `xml:"dataCenterInfo" json:"dataCenterInfo"`
	LeaseInfo                     *LeaseInfo             `xml:"leaseInfo" json:"leaseInfo"`
	Metadata                      map[string]interface{} `xml:"metadata" json:"metadata"`
	IsCoordinatingDiscoveryServer interface{}            `xml:"isCoordinatingDiscoveryServer" json:"isCoordinatingDiscoveryServer"`
}

type Port struct {
	Port   int         `json:"$"`
	Enable interface{} `json:"@enabled"`
}

type DataCenterInfo struct {
	Name  string `json:"name" xml:"name"`
	Class string `json:"@class"`
}

type LeaseInfo struct {
	RenewalIntervalInSecs int   `xml:"renewalIntervalInSecs" json:"renewalIntervalInSecs"`
	DurationInSecs        int   `xml:"durationInSecs" json:"durationInSecs"`
	RegistrationTimestamp int64 `xml:"registrationTimestamp" json:"registrationTimestamp"`
	LastRenewalTimestamp  int64 `xml:"lastRenewalTimestamp" json:"lastRenewalTimestamp"`
	RenewalTimestamp      int64 `xml:"renewalTimestamp" json:"renewalTimestamp"`
	EvictionTimestamp     int64 `xml:"evictionTimestamp" json:"evictionTimestamp"`
	ServiceUpTimestamp    int64 `xml:"serviceUpTimestamp" json:"serviceUpTimestamp"`
}

func (i *Instance) Id() string {
	if i.InstanceId != "" {
		return i.InstanceId
	}
	return i.HostName + ":" + i.App + ":" + strconv.Itoa(i.Port.Port)
}

func (i *Instance) Init() {
	// InstanceId
	i.InstanceId = i.Id()
	i.VipAddress = i.App
	i.SecureVipAddress = i.App
	// LeaseInfo
	leaseInfo := i.LeaseInfo
	if leaseInfo == nil {
		leaseInfo = &LeaseInfo{}
	}
	if leaseInfo.RenewalIntervalInSecs == 0 {
		leaseInfo.RenewalIntervalInSecs = 30
	}
	if leaseInfo.DurationInSecs == 0 {
		leaseInfo.DurationInSecs = 90
	}
	i.LeaseInfo = leaseInfo
	// SecurePort
	if i.SecurePort == nil {
		i.SecurePort = &Port{
			Port:   443,
			Enable: "false",
		}
	}
	// MetaData
	if i.Metadata == nil {
		i.Metadata = map[string]interface{}{"@class": "java.util.Collections$EmptyMap"}
	}
}
