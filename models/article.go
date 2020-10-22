package models

import (
	"time"
)

type Article struct {
	ID      int       `json:"id,omitempty"`
	Cid     int       `json:"cid,omitempty"`
	Title   string    `json:"title,omitempty"`
	Author  string    `json:"author,omitempty"`
	Content string    `json:"content,omitempty"`
	Hits    string    `json:"hits,omitempty"`
	Ctime   time.Time `json:"ctime,omitempty"`
}
