package eureka

import (
	"fmt"
	"testing"
)

func TestRegisterApp(t *testing.T) {
	ip := "118.178.230.252"
	ins := Instance{
		HostName:         ip,
		App:              "goaaa",
		Port:             &Port{Port: 80, Enable: true},
		IPAddr:           ip,
		VipAddress:       ip,
		SecureVipAddress: ip,
		HealthCheckUrl:   "http://" + ip + "/health",
		StatusPageUrl:    "http://" + ip + "/status",
		HomePageUrl:      "http://" + ip,
		Status:           "UP",
		DataCenterInfo:   &DataCenterInfo{Name: "MyOwn"},
	}

	serverUrls := []string{""}
	e := Eureka{ServiceUrls: serverUrls}
	err := e.RegisterInstane(&ins)
	if err != nil {
		fmt.Println(err)
		t.Fatal()
	}
	t.Log("register service success.")
}
