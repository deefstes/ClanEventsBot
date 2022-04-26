package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

func JSONError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	j, _ := json.Marshal(err)
	w.Write(j)
	// json.NewEncoder(w).Encode(err)
}

func middlewareContentType(nextHandler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		apikey := r.Header["X-Api-Key"]
		if len(apikey) > 0 && apikey[0] != config.ApiKey {
			JSONError(w, ErrorMessage{Message: "unauthorised"}, http.StatusUnauthorized)
			return
		}
		nextHandler(w, r)
	}
}

func catchAllHandler(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String()

	rsp := CatchAllResponse{
		Method:        r.Method,
		RequestUri:    r.RequestURI,
		Url:           r.URL,
		ContentLength: r.ContentLength,
		Host:          r.Host,
		Proto:         r.Proto,
		RemoteAddr:    r.RemoteAddr,
		Body:          body,
	}
	j, _ := json.Marshal(rsp)
	w.Write(j)
}

// GET /api/health
func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		JSONError(
			w,
			UnsupportedResponse{SupportedMethods: []string{"GET"}},
			http.StatusMethodNotAllowed,
		)
		return
	}

	status := "ok"
	var info string
	d, err := db.Ping()
	if err != nil {
		status = "degraded"
		info = err.Error()
	}

	rsp := HealthResponse{
		Status:     status,
		Info:       info,
		DBResponse: d.String(),
		LiveTime:   liveTime,
		DebugLevel: config.DebugLevel,
	}
	j, _ := json.Marshal(rsp)
	w.Write(j)
}

// GET /api/events?guildID={Guild_ID}
func listEventsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		JSONError(
			w,
			UnsupportedResponse{SupportedMethods: []string{"GET"}},
			http.StatusMethodNotAllowed,
		)
		return
	}

	guilID := r.URL.Query().Get("guildId")
	if guilID == "" {
		JSONError(w, ErrorMessage{Message: "guildId not provided"}, http.StatusBadRequest)
		return
	}

	events, err := db.GetEvents(guilID, "all", time.Time{})
	if err != nil {
		JSONError(w, ErrorMessage{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	rsp, err := json.Marshal(events)
	if err != nil {
		JSONError(w, ErrorMessage{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(rsp))
}
