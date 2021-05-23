package puppetweb

type BaseRequest struct {
	DeviceID string `json:"DeviceID"`
	Sid      string `json:"Sid"`
	Skey     string `json:"Skey"`
	Uin      string `json:"Uin"`
}

type BaseResponse struct {
	Ret    int    `json:"Ret"`
	ErrMsg string `json:"ErrMsg"`
}
