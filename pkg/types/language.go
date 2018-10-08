package types

import "time"

type LanguageDetectionLogEntry struct {
	Id        string    `json:"id"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Type      string    `json:"type"`
}