package entity

// Thing is db record of thing
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"name"`
	LastName  string `json:"last"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
}
