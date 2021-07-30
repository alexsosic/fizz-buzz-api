package models

// Stats is a statistics model
type Stats struct {
	Request string `json:"request" gorm:"primary_key"`
	Hits    int    `json:"hits"`
}
