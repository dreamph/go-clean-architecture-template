package models

type Bytes struct {
	Data []byte `json:"data"`
}

type KV[K, V any] struct {
	Key   K `json:"key"`
	Value V `json:"value"`
}
