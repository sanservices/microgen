package entity

// ThingRec is db record of thing
type ThingRec struct {
	ID         uint `json:"id" db:"ID" fake:"{uint8}"`
	CategoryID uint `json:"category_id" db:"CategoryID" fake:"{uint8}"`
}
