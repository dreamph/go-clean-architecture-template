package models

type CertDescription struct {
	RefID       string `json:"refId,omitempty"`
	RefNo       string `json:"refNo,omitempty"`
	RefFileHash string `json:"refFileHash,omitempty"`
	Desc        string `json:"desc,omitempty"`
	Ref1        string `json:"ref1,omitempty"`
	Ref2        string `json:"ref2,omitempty"`
	Ref3        string `json:"ref3,omitempty"`
}
