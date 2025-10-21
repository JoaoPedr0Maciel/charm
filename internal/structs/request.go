package structs

import (
	"net/http"
	"time"
)

type HTTPMethod struct {
	Name           string
	HasBody        bool
	HTTPFunc       func(string, string, string, string, string) (*http.Response, error)
	HTTPFuncNoBody func(string, string, string, string) (*http.Response, error)
}

type RequestOptions struct {
	Method      string
	URL         string
	Bearer      string
	Basic       string
	ContentType string
	Data        string
}

type Display struct {
	Method      string
	URL         string
	Bearer      string
	Basic       string
	ContentType string
	Data        string
	Request     *http.Request
	Response    *http.Response
	Body        []byte
	AuthHeader  string
	TotalTime   time.Duration
	Timing      *TimingInfo
}

func NewDisplay(method, url string) *Display {
	return &Display{
		Method: method,
		URL:    url,
	}
}

func (d *Display) WithAuth(bearer, basic, authHeader string) *Display {
	d.Bearer = bearer
	d.Basic = basic
	d.AuthHeader = authHeader
	return d
}

func (d *Display) WithContent(contentType, data string) *Display {
	d.ContentType = contentType
	d.Data = data
	return d
}

func (d *Display) WithHTTP(req *http.Request, resp *http.Response, body []byte) *Display {
	d.Request = req
	d.Response = resp
	d.Body = body
	return d
}

func (d *Display) WithTiming(totalTime time.Duration, timing *TimingInfo) *Display {
	d.TotalTime = totalTime
	d.Timing = timing
	return d
}
