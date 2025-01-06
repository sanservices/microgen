package entity

// User is db record of user
type User struct {
	ID        uint32 `json:"id"`
	FirstName string `json:"name"`
	LastName  string `json:"last"`
	Age       uint32 `json:"age"`
	Email     string `json:"email"`
}
