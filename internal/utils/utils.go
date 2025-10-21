package utils

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/http/httptrace"
	"net/url"
	"strings"
	"time"

	"github.com/JoaoPedr0Maciel/charm/internal/structs"
	"github.com/JoaoPedr0Maciel/charm/internal/ui"
)

const (
	DefaultContentType = "application/json"
	BearerPrefix       = "Bearer "
	BasicPrefix        = "Basic "
)

func DoRequest(opts structs.RequestOptions) (*http.Response, error) {
	if err := validateURL(opts.URL); err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	startTime := time.Now()
	timing := &structs.TimingInfo{}

	req, err := createRequest(opts, timing)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	authHeader := addAuthentication(req, opts.Bearer, opts.Basic)
	setContentType(req, opts.ContentType, opts.Data)

	timing.RequestStart = time.Now()
	resp, err := http.DefaultClient.Do(req)
	timing.RequestDone = time.Now()

	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	body, err := readResponseBody(resp)
	timing.ResponseDone = time.Now()
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	totalTime := timing.ResponseDone.Sub(startTime)

	display := structs.NewDisplay(opts.Method, opts.URL).
		WithAuth(opts.Bearer, opts.Basic, authHeader).
		WithContent(opts.ContentType, opts.Data).
		WithHTTP(req, resp, body).
		WithTiming(totalTime, timing)

	ui.Display(*display)

	return resp, nil
}

func validateURL(rawURL string) error {
	if rawURL == "" {
		return fmt.Errorf("URL cannot be empty")
	}

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return err
	}

	if parsedURL.Scheme == "" {
		return fmt.Errorf("URL must have a scheme (http:// or https://)")
	}

	if parsedURL.Host == "" {
		return fmt.Errorf("URL must have a host")
	}

	return nil
}

func createRequest(opts structs.RequestOptions, timing *structs.TimingInfo) (*http.Request, error) {
	var bodyReader io.Reader
	if opts.Data != "" {
		bodyReader = strings.NewReader(opts.Data)
	}

	req, err := http.NewRequest(opts.Method, opts.URL, bodyReader)
	if err != nil {
		return nil, err
	}

	trace := createClientTrace(timing)
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	return req, nil
}

func createClientTrace(timing *structs.TimingInfo) *httptrace.ClientTrace {
	return &httptrace.ClientTrace{
		DNSStart: func(_ httptrace.DNSStartInfo) {
			timing.DNSStart = time.Now()
		},
		DNSDone: func(_ httptrace.DNSDoneInfo) {
			timing.DNSDone = time.Now()
		},
		ConnectStart: func(_, _ string) {
			timing.ConnectStart = time.Now()
		},
		ConnectDone: func(_, _ string, _ error) {
			timing.ConnectDone = time.Now()
		},
		TLSHandshakeStart: func() {
			timing.TLSStart = time.Now()
		},
		TLSHandshakeDone: func(_ tls.ConnectionState, _ error) {
			timing.TLSDone = time.Now()
		},
		GotFirstResponseByte: func() {
			timing.ResponseStart = time.Now()
		},
	}
}

func addAuthentication(req *http.Request, bearer, basic string) string {
	if bearer != "" {
		authHeader := BearerPrefix + bearer
		req.Header.Set("Authorization", authHeader)
		return authHeader
	}

	if basic != "" {
		encoded := base64.StdEncoding.EncodeToString([]byte(basic))
		authHeader := BasicPrefix + encoded
		req.Header.Set("Authorization", authHeader)
		return authHeader
	}

	return ""
}

func setContentType(req *http.Request, contentType, data string) {
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
		return
	}

	if data != "" {
		req.Header.Set("Content-Type", DefaultContentType)
	}
}

func readResponseBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
