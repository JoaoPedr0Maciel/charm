package structs

import "time"

type Response struct {
	StatusCode    int
	Url           string
	Method        string
	Authorization *string
	ContentType   string
	Body          *[]byte
}

type TimingInfo struct {
	DNSStart      time.Time
	DNSDone       time.Time
	ConnectStart  time.Time
	ConnectDone   time.Time
	TLSStart      time.Time
	TLSDone       time.Time
	RequestStart  time.Time
	RequestDone   time.Time
	ResponseStart time.Time
	ResponseDone  time.Time
}
