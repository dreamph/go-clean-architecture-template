package models

type AuthKey struct {
	AppID string `json:"appId"`
	//Code  string `json:"code"`
	//AuthType int32  `json:"authType"`

	//Enc
	Token  string `json:"token"`
	Secret string `json:"secret"`
	Key    string `json:"key"`
	Iv     string `json:"iv"`
}
