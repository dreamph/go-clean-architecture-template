package models

type ClientResponse[DT any, ET any] struct {
	Result      *DT
	ErrorResult *ET
}
