package entity

// User is DB ID FirstName LastName
type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// Todos is DB ID Text Status
type Todos struct {
	ID     int
	Text   string
	Status string
}
