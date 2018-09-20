package eureka

import (
	"testing"
)

func TestRegisterApp(t *testing.T) {
	ip := "localhost"
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
