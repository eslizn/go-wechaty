package puppetweb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type BaseRequest struct {
	DeviceID string `json:"DeviceID"`
	Sid      string `json:"Sid"`
	Skey     string `json:"Skey"`
	Uin      string `json:"Uin"`
}

func (p *PuppetWeb) init() error {
	buff, err := json.Marshal(struct {
		BaseRequest interface{} `json:"BaseRequest"`
	}{BaseRequest: struct {
		DeviceID string `json:"DeviceID"`
		Sid      string `json:"Sid"`
		Skey     string `json:"Skey"`
		Uin      string `json:"Uin"`
	}{
		DeviceID: p.option.DeviceID,
		Sid:      p.ticket.Wxsid,
		Skey:     p.ticket.Skey,
		Uin:      p.ticket.Wxuin,
	},
	})
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", (&url.URL{
		Scheme: "https",
		Host:   "wx.qq.com",
		Path:   "/cgi-bin/mmwebwx-bin/webwxinit",
		RawQuery: (url.Values{
			"pass_ticket": []string{p.ticket.PassTicket},
			"skey":        []string{p.ticket.Skey},
			"r":           []string{fmt.Sprintf("%d", time.Now().Unix())},
		}).Encode(),
	}).String(), bytes.NewReader(buff))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	rsp, err := p.client.Do(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()
	buff, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}
	return nil
}

func (p *PuppetWeb) statusNotify() {

}

func (p *PuppetWeb) syncCheck() {

}

func (p *PuppetWeb) sync() {

}
