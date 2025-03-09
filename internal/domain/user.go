package domain

type User struct {
	Id   int64
	tgId int64
	tgName string
	timeLastActivity int64

	age int16
	geo string
	gender string
	genderSearch string
	photoUrl string
	description string
}

