package entity

type Filter struct {
	ID     int        `json:"id"`
	Name   string     `json:"name"`
	Hint   *string    `json:"hint"`
	Type   FilterType `json:"type"`
	Values []string   `json:"values"`
}
