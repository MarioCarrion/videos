package main

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
	"time"
)

const (
	PriorityNone   Priority = iota // PriorityNone indicates the task needs to be prioritized.
	PriorityLow                    // PriorityLow indicates a non urgent task.
	PriorityMedium                 // PriorityMedium indicates a task that should be completed soon.
	PriorityHigh                   // PriorityHigh indicates an urgent task that must be completed as soon as possible.
)

type (
	Task struct {
		Completed   bool     `json:"completed"`
		Priority    Priority `json:"priority"`
		ID          int64    `json:"id,string"`
		Description string   `json:"description,omitempty"`
		Dates       Dates    `json:",inline"` // "inline": new in Go 1.25
	}

	Priority int8

	Dates struct {
		Start time.Time `json:"startDate,format:'2006-01-02'"` // "format:" new in Go 1.25
		Due   time.Time `json:"dueDate,format:'2006-01-02'"`
	}
)

func main() {
	newYork, _ := time.LoadLocation("America/New_York")

	task := Task{
		Completed:   false,
		Priority:    PriorityMedium,
		ID:          1_234_567,
		Description: "", // will be omitted due to "omitempty"
		Dates: Dates{
			Start: time.Date(2025, time.August, 29, 0, 0, 0, 0, newYork),
			Due:   time.Date(2025, time.August, 30, 0, 0, 0, 0, newYork),
		},
	}

	// Custom marshaler for Priority type without implementing json.Marshaler
	// Allows external types to be marshaled without modifying their source code.

	priorityMarshaler := json.MarshalToFunc(
		func(enc *jsontext.Encoder, val Priority) error {
			switch val {
			case PriorityLow:
				return enc.WriteToken(jsontext.String("low"))
			case PriorityMedium:
				return enc.WriteToken(jsontext.String("medium"))
			case PriorityHigh:
				return enc.WriteToken(jsontext.String("high"))
			}
			return enc.WriteToken(jsontext.String("none"))
		},
	)

	//-

	b, _ := json.Marshal(task, json.WithMarshalers(priorityMarshaler))
	fmt.Println(string(b))
}
