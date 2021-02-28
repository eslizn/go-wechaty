package puppetweb

import (
	"encoding/xml"
	"math/rand"
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

func (p *PuppetWeb) mockDeviceID() string {
	sets := []byte("0123456789qwertyuiopasdfghjklzxcvbnm")
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 15; i++ {
		result = append(result, sets[r.Intn(len(sets))])
	}
	return "e" + string(result)
}
