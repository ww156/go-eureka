package eureka

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/json-iterator/go"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Eureka struct {
	ServiceUrls []string
	Client      *http.Client
	Json        bool
}

func NewEureka(serverUrls []string, client *http.Client) (*Eureka, error) {
	if len(serverUrls) == 0 {
		return nil, errors.New("missing eureka url.")
	}
	if client == nil {
		client = http.DefaultClient
	}
	eureka := &Eureka{
		ServiceUrls: serverUrls,
		Client:      client,
	}
	return eureka, nil
}

func (e *Eureka) pickServerUrl() string {
	urls := e.ServiceUrls
	l := len(urls)
	if l == 0 {
		return ""
		//panic(errors.New("no valid eureka server."))
	}
	if l == 1 {
		if checkIp(urls[0] + "/apps") {
			return urls[0]
		}
		return ""
		//panic(errors.New("no valid eureka server."))
	}
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(l)
	url := urls[r]
	if checkIp(url + "/apps") {
		return url
	} else {
		for i := 0; i < l; i++ {
			if checkIp(urls[i] + "/apps") {
				return urls[i]
			}
		}
	}
	return ""
	//panic(errors.New("no valid eureka server."))
}

// 注册实例
func (e *Eureka) RegisterInstane(i *Instance) error {
	urls := e.ServiceUrls
	i.Init()
	// Instance数据构建
	app := App{
		Instance: i,
	}
	data, err := json.Marshal(&app)
	if err != nil {
		return err
	}

	for _, url := range urls {
		req, err := http.NewRequest("POST", url+"/apps/"+i.App, bytes.NewReader(data))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			fmt.Println(err)
			continue
		}
		res, err := e.Client.Do(req)
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer res.Body.Close()
		if res.StatusCode != 204 {
			fmt.Println("server " + url + " is't" + " registed")
		}
	}

	return nil
}

// 发送心跳
func (e *Eureka) SendHeartBeat(i *Instance, duration time.Duration) {
	go func() {
		ticker := time.NewTicker(duration)
		urls := e.ServiceUrls
		if len(urls) == 0 {
			fmt.Println("missing eureka url.")
		}
		for {
			select {
			case <-ticker.C:
				defer func() {
					recover()
					fmt.Println("SendHeartBeat Error.")
				}()
				l := len(urls)
				rand.Seed(time.Now().UnixNano())
				n := rand.Intn(l)
				req, err := http.NewRequest("PUT", urls[n]+"/apps/"+i.App+"/"+i.InstanceId, nil)
				if err != nil {
					fmt.Println(err)
					continue
				}
				res, err := e.Client.Do(req)
				if err != nil {
					fmt.Println(err)
					continue
				}
				res.Body.Close()
				statusCode := res.StatusCode
				if statusCode != 200 {
					if statusCode == 404 {
						e.RegisterInstane(i)
					} else {
						fmt.Println(errors.New("unknown error."))
					}
				}
			}
		}
	}()
}

// 获取APP
func (e *Eureka) GetApp(appid string) (*Application, error) {
	url := e.pickServerUrl()
	//fmt.Println("GET", url+"/"+appid)
	req, err := http.NewRequest("GET", url+"/apps/"+appid, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, err := e.Client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//fmt.Println(string(body))
	resp.Body.Close()
	result := Application{}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("response StatusCode is not 200，but " + strconv.Itoa(resp.StatusCode))
	}
	err = jsoniter.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// 获取APP url列表
func (e *Eureka) GetAppUrls(appid string) []string {
	app, err := e.GetApp(appid)
	n := 0
	for err != nil {
		n += 1
		time.Sleep(time.Millisecond * 200)
		app, err = e.GetApp(appid)
		if n > 10 {
			break
		}
	}
	if err != nil {
		return []string{}
	}
	urls := []string{}
	for _, ins := range app.Application.Instance {
		if ins.Status == "UP" {
			url := ins.IPAddr + ":" + strconv.Itoa(ins.Port.Port)
			if checkIp(ins.HealthCheckUrl) {
				urls = append(urls, url)
			}
		}
	}
	if len(urls) == 0 {
		return []string{}
	}
	return urls
}

// 删除实例
func (e *Eureka) DelInstance(appid, instanceid string) error {
	urls := e.ServiceUrls
	if len(urls) == 0 {
		return errors.New("missing eureka url.")
	}
	req, err := http.NewRequest("DELETE", urls[0]+"/"+appid+"/"+instanceid, nil)
	if err != nil {
		return err
	}
	res, err := e.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		return nil
	} else {
		return errors.New("unknown error.")
	}
}
