package data

type RequestStatus uint
const (
	INPROCESS RequestStatus = iota
	ACCEPTED
	DENIED
)
