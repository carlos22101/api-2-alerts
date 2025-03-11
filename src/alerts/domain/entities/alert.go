package entities

import "time"

type Alert struct {
	ID          int       `json:"id"`
	SensorID    string    `json:"sensorId"`
	EventTime   string    `json:"event_timestamp"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
