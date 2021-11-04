package entity

// Thing is db record of thing
type Thing struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
