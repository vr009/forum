package models

type Forum struct {
	Title   string `json:"title"`
	User    string `json:"user"`
	Slug    string `json:"slug"`
	Posts   int32  `json:"posts"`
	Threads int32  `json:"threads"`
}
