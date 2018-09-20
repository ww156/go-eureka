package eureka

import (
	"bytes"
	"encoding/json"
	"errors"
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

// 注册实例
func (e *Eureka) RegisterInstane(instance *Instance) error {
	urls := e.ServiceUrls
	if len(urls) == 0 {
		return errors.New("missing eureka url.")
	}
	instance.Init()
	// Instance数据构建
	app := App{
		Instance: instance,
	}
	data, err := json.Marshal(&app)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", urls[0]+"/apps/"+instance.App, bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}
	res, err := e.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 204 {
		return errors.New(strconv.Itoa(res.StatusCode))
	}
	return nil
}

// 发送心跳
func (e *Eureka) SendHeartBeat(appid, instanceid string, duration time.Duration) {
	go func() {
		ticker := time.NewTicker(duration)
		urls := e.ServiceUrls
		if len(urls) == 0 {
			panic("missing eureka url.")
		}
		for {
			select {
			case <-ticker.C:
				req, err := http.NewRequest("PUT", urls[0]+"/"+appid+"/"+instanceid, nil)
				if err != nil {
					panic(err)
				}
				res, err := e.Client.Do(req)
				if err != nil {
					panic(err)
				}
				defer res.Body.Close()
				if res.StatusCode == 404 {
					panic(errors.New("instanceID doesn’t exist."))
				} else if res.StatusCode != 200 {
					panic(errors.New("unknown error."))
				}
			}
		}
	}()
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
