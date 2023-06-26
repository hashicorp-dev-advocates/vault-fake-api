package main

import "time"

type EventInstance struct {
	Id        string    `json:"id"`
	Source    string    `json:"source"`
	Type      string    `json:"type"`
	Data      Data      `json:"data"`
	Time      time.Time `json:"time"`
	EventType string    `json:"event_type"`
}

type Data struct {
	Event Event `json:"event"`
}

type Event struct {
	Id       string   `json:"id"`
	Metadata Metadata `json:"metadata"`
}

type Metadata struct {
	SourceIpAddress string `json:"source_ip_address"`
	Username        string `json:"username"`
}
