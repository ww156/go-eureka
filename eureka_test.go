package eureka

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestRegisterApp(t *testing.T) {
	app := "GOAAA"
	ip := "localhost"
	port := 80
	ins := Instance{
		HostName:         ip,
		App:              app,
		Port:             &Port{Port: port, Enable: "true"},
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

	serverUrls := []string{"http://192.168.1.19:1111/eureka"}
	e, err := NewEureka(serverUrls, nil)
	if err != nil {
		t.Fatal(err)
	}
	err = e.RegisterInstane(&ins)
	if err != nil {
		t.Fatal(err)
	}
	e.SendHeartBeat(&ins, time.Second*20)
	t.Log("register service success.")
	for {
	}
}

func TestGetApp(t *testing.T) {
	serverUrls := []string{"http://192.168.1.19:1111/eureka"}
	e, err := NewEureka(serverUrls, nil)
	if err != nil {
		t.Fatal(err)
	}
	urls := e.GetAppUrls("GOAAA")
	fmt.Println(urls)
}
