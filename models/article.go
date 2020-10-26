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

type Page struct {
	PageNo   int `json:"page_no,omitempty"`
	PageSize int `json:"page_size,omitempty"`
	ClassID  int `json:"class_id,omitempty"`
}
