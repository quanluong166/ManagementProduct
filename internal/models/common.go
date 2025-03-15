package models

type Paging struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
