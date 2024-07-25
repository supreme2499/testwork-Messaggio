package database

// модель записи в бд
type Contents struct {
	Content string `json:"content"`
	Status  string `json:"status"`
}
