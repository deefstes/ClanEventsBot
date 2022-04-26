package main

import (
	"net/url"
	"time"
)

type UnsupportedResponse struct {
	SupportedMethods []string `json:"supportedMethods,omitempty"`
}

type CatchAllResponse struct {
	Method        string   `json:"method,omitempty"`
	RequestUri    string   `json:"requestUri,omitempty"`
	Url           *url.URL `json:"url,omitempty"`
	ContentLength int64    `json:"contentLength,omitempty"`
	Host          string   `json:"host,omitempty"`
	Proto         string   `json:"proto,omitempty"`
	RemoteAddr    string   `json:"remoteAddr,omitempty"`
	Body          string   `json:"body,omitempty"`
}

type HealthResponse struct {
	Status     string    `json:"status,omitempty"`
	Info       string    `json:"info,omitempty"`
	DBResponse string    `json:"databaseResponse,omitempty"`
	LiveTime   time.Time `json:"liveTime,omitempty"`
	DebugLevel int       `json:"debugLevel,omitempty"`
}
