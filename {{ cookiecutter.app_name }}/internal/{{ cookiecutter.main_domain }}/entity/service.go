package entity

// Thing is db record of thing
type Thing struct {
	ID         uint   `json:"id"`
	CategoryID uint   `json:"category_id"`
	Image      string `json:"image"`
}
