package types

type Show struct {
	MALID       int    `json:"mal_id" db:"mal_id"` // Use as primary key
	Title       string `json:"title" db:"title"`
	Source      string `json:"source" db:"source"`
	SourceID    string `json:"source_id" db:"source_id"`
	MainPicture string `json:"main_picture" db:"main_picture"`
}
