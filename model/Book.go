package model

type Book struct {
	Link    string   `json:"link"`
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
	Cover   string   `json:"cover"`
	Rating  string   `json:"rating"`
}

func (b Book) IsEmpty() bool {
	return b.Link == ""
}
