package puppetweb

import (
	"testing"
	"time"
)

func TestPuppetWeb(t *testing.T) {
	p, err := NewPuppetWeb(nil)
	if err != nil {
		t.Error(err)
		return
	}
	err = p.getUUID()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(p.getQRCodeUrl())
	//png, err := qrcode.Encode(p.qrCodeUrl(), qrcode.Medium, 256)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//t.Log(png)
	time.Sleep(time.Second)
	err = p.login("1")
	if err != nil {
		t.Error(err)
		return
	}
	time.Sleep(time.Second)
	err = p.login("0")
	if err != nil {
		t.Error(err)
		return
	}
	time.Sleep(time.Second)
	err = p.newLoginPage()
	if err != nil {
		t.Error(err)
		return
	}
	time.Sleep(time.Second)
	resp, err := p.init()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", resp)
}
