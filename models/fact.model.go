package models

type Fact struct {
	ID       int    `json:"id,omitempty"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
