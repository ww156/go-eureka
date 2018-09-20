package eureka

import (
	"strconv"
	"testing"
)

func TestRegisterApp(t *testing.T) {
	ip := "localhost"
	port := 80
	ins := Instance{
		HostName:         ip,
		App:              "goaaa",
		Port:             &Port{Port: port, Enable: true},
		IPAddr:           ip,
		VipAddress:       ip,
		SecureVipAddress: ip,
		HealthCheckUrl:   "http://" + ip + ":" + strconv.Itoa(port) + "/health",
		StatusPageUrl:    "http://" + ip + ":" + strconv.Itoa(port) + "/status",
		HomePageUrl:      "http://" + ip + ":" + strconv.Itoa(port),
		Status:           "UP",
		DataCenterInfo: &DataCenterInfo{
			Name:  "MyOwn",
			Class: "com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo",
		},
	}

	serverUrls := []string{"http://localhost:8761/eureka"}
	e, err := NewEureka(serverUrls, nil)
	if err != nil {
		t.Fatal(err)
	}
	err = e.RegisterInstane(&ins)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("register service success.")
}
