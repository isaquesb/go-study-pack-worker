package domain

type Pack struct {
	ID string `json:"id"`
	Date   string `json:"date"`
	Owner  Owner  `json:"owner,omitempty"`
	Target Target `json:"target"`
}
