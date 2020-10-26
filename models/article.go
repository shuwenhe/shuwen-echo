package models

import (
	"time"
)

type Article struct {
	ID      int       `json:"id,omitempty"`
	Cid     int       `json:"cid,omitempty"`
	Class   *Class    `json:"class,omitempty"`
	Title   string    `json:"title,omitempty"`
	Author  string    `json:"author,omitempty"`
	Content string    `json:"content,omitempty"`
	Hits    int       `json:"hits,omitempty"`
	Ctime   time.Time `json:"ctime,omitempty"`
}
