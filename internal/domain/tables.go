package domain

type MatchTable struct {
	ID         int
	UserID     int
	WantUserID int
	Status     string
	Message    string
}

type SuggestionUser struct {
	ID           int
	UserID       int
	SuggestionID int
}

type Suggestion struct {
	ID                 int
	ActivityCategoryID int
	Name               string
	Description        string
	PhotoURLs          []string
}

type SuggestionCategory struct {
	ID   int
	Name string
}
