package database

// модель записи в бд
type Contents struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Status  string `json:"status"`
}
