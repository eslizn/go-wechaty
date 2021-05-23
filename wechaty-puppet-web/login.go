package puppetweb

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

type xmlTicket struct {
	XMLName     xml.Name `xml:"error"`
	Ret         int      `xml:"ret"`
	Message     string   `xml:"message"`
	Skey        string   `xml:"skey"`
	Wxsid       string   `xml:"wxsid"`
	Wxuin       string   `xml:"wxuin"`
	PassTicket  string   `xml:"pass_ticket"`
	IsGrayscale int      `xml:"isgrayscale"`
}

func (p *PuppetWeb) makeDeviceID() string {
	sets := []byte("0123456789qwertyuiopasdfghjklzxcvbnm")
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 15; i++ {
		result = append(result, sets[r.Intn(len(sets))])
	}
	return "e" + string(result)
}

// getUUID get uuid from new login page
func (p *PuppetWeb) getUUID() error {
	req, err := http.NewRequest("GET", (&url.URL{
		Scheme: "https",
		Host:   "login.weixin.qq.com",
		Path:   "/jslogin",
		RawQuery: (url.Values{
			"appid":        []string{p.option.AppId},
			"fun":          []string{"new"},
			"lang":         []string{p.option.Lang},
			"redirect_uri": []string{"https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxnewloginpage"},
			"_":            []string{fmt.Sprintf("%d", time.Now().Unix())},
		}).Encode(),
	}).String(), nil)
	if err != nil {
		return err
	}
	rsp, err := p.client.Do(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()
	buff, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}
	regex, err := regexp.Compile("^window.QRLogin.code = ([0-9]+); window.QRLogin.uuid = \"(.*)\";$")
	if err != nil {
		return err
	}
	params := regex.FindStringSubmatch(string(buff))
	if params[1] != "200" {
		return errors.New(string(buff))
	}
	p.uuid = params[2]
	return nil
}

// getQRCodeUrl get qrcode url
func (p *PuppetWeb) getQRCodeUrl() string {
	return (&url.URL{
		Scheme: "https",
		Host:   "login.weixin.qq.com",
		Path:   "/qrcode/" + p.uuid,
	}).String()
}

// getQRCode get qrcode data
func (p *PuppetWeb) getQRCode() ([]byte, error) {
	req, err := http.NewRequest("GET", p.getQRCodeUrl(), nil)
	if err != nil {
		return nil, err
	}
	rsp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	return ioutil.ReadAll(rsp.Body)
}

// login
func (p *PuppetWeb) login(tip string) error { //tip 1:未扫描 0:已扫描
	req, err := http.NewRequest("GET", (&url.URL{
		Scheme: "https",
		Host:   "login.weixin.qq.com",
		Path:   "/cgi-bin/mmwebwx-bin/login",
		RawQuery: (url.Values{
			"loginicon": []string{"true"},
			"uuid":      []string{p.uuid},
			"tip":       []string{tip},
			"r":         []string{fmt.Sprintf("%d", time.Now().Unix())},
			"_":         []string{fmt.Sprintf("%d", time.Now().Unix())},
		}).Encode(),
	}).String(), nil)
	if err != nil {
		return err
	}
	rsp, err := p.client.Do(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()
	buff, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}
	regex, err := regexp.Compile("window.code=([0-9]+)")
	if err != nil {
		return err
	}
	params := regex.FindStringSubmatch(string(buff))
	if len(params) < 2 {
		return errors.New("not found window.code in response")
	}
	code, err := strconv.Atoi(params[1])
	if err != nil {
		return err
	}
	//@todo code = 408 timeout
	if code == 200 {
		regex, err := regexp.Compile("window.redirect_uri=\"(.*)\";")
		if err != nil {
			return err
		}
		params := regex.FindStringSubmatch(string(buff))
		if len(params) < 2 {
			return errors.New("not found window.redirect_uri in response")
		}
		p.redirectUri = params[1]
	}
	return nil
}

//get login params
func (p *PuppetWeb) newLoginPage() error {
	parse, err := url.Parse(p.redirectUri)
	if err != nil {
		return err
	}
	query := parse.Query()
	query.Set("fun", "new")
	//query.Set("version", "v2")
	parse.RawQuery = query.Encode()
	req, err := http.NewRequest("GET", parse.String(), nil)
	if err != nil {
		return err
	}
	rsp, err := p.client.Do(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()
	buff, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}
	if err := xml.Unmarshal(buff, &p.ticket); err != nil {
		return err
	}
	if p.ticket.Ret != 0 {
		return fmt.Errorf("[%d]%s", p.ticket.Ret, p.ticket.Message)
	}
	return nil
}
