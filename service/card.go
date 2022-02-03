package service

type Card struct {
	Value string `json:"value,omitempty"`
	Suit  string `json:"suit,omitempty"`
	Code  string `json:"code,omitempty"`
}
