package data

type QueryType uint

const (
	PROFILE QueryType = iota
	CONTENT
	BOTH
)