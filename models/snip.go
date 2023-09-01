package models

import "time"

type Snip struct {
	SnipId      int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DateTime    time.Time `json:"datetime"`
}
