package main

import (
	"net/url"
	"time"
)

type unsupportedResponse struct {
	SupportedMethods []string `json:"supportedMethods,omitempty"`
}

type catchAllResponse struct {
	Method        string   `json:"method,omitempty"`
	RequestURI    string   `json:"requestUri,omitempty"`
	URL           *url.URL `json:"url,omitempty"`
	ContentLength int64    `json:"contentLength,omitempty"`
	Host          string   `json:"host,omitempty"`
	Proto         string   `json:"proto,omitempty"`
	RemoteAddr    string   `json:"remoteAddr,omitempty"`
	Body          string   `json:"body,omitempty"`
}

type healthResponse struct {
	Status     string    `json:"status,omitempty"`
	Info       string    `json:"info,omitempty"`
	DBResponse string    `json:"databaseResponse,omitempty"`
	LiveTime   time.Time `json:"liveTime,omitempty"`
	DebugLevel int       `json:"debugLevel,omitempty"`
}
