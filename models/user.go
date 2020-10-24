package models

type User struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"num,omitempty"`
	Password string `json:"password,omitempty"`
	Phone    string `json:"phone,omitempty"`
}
