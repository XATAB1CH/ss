package domain

type User struct {
	ID               int64
	TgID             int64
	TgName           string
	TimeLastActivity int64

	Age          int16
	Geo          string
	Gender       string
	GenderSearch string
	PhotoUrl     string
	Description  string
}
