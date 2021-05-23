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

type Sync struct {
	Key int `json:"Key"`
	Val int `json:"Val"`
}

type SyncKey struct {
	Count int    `json:"Count"`
	List  []Sync `json:"List"`
}

type InitResponse struct {
	BaseResponse        BaseResponse `json:"BaseResponse"`
	Count               int          `json:"Count"`
	ContactList         []Contact    `json:"ContactList"`
	SyncKey             SyncKey      `json:"SyncKey"`
	User                Contact      `json:"User"`
	ChatSet             string       `json:"ChatSet"`
	SKey                string       `json:"SKey"`
	ClientVersion       int          `json:"ClientVersion"`
	SystemTime          int          `json:"SystemTime"`
	GrayScale           int          `json:"GrayScale"`
	InviteStartCount    int          `json:"InviteStartCount"`
	MPSubscribeMsgCount int          `json:"MPSubscribeMsgCount"`
	MPSubscribeMsgList  []Subscribe  `json:"MPSubscribeMsgList"`
	ClickReportInterval int          `json:"ClickReportInterval"`
}

func (p *PuppetWeb) init() (*InitResponse, error) {
	buff, err := json.Marshal(struct {
		BaseRequest BaseRequest `json:"BaseRequest"`
	}{BaseRequest: BaseRequest{
		DeviceID: p.option.DeviceID,
		Sid:      p.ticket.Wxsid,
		Skey:     p.ticket.Skey,
		Uin:      p.ticket.Wxuin,
	}})
	if err != nil {
		return nil, err
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
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	rsp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	buff, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	result := &InitResponse{}
	return result, json.Unmarshal(buff, result)
}

func (p *PuppetWeb) notify() {

}

func (p *PuppetWeb) check() {

}

func (p *PuppetWeb) sync() {

}
