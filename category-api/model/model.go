package model

type Category struct {
	ID        uint64  `json:"id"`
	Category  string  `json:"category"`
	CreatedAt string  `json:"created_at"`
	UpdateAt  string  `json:"updated_at"`
	DeletedAt *string `json:"deleted_at,omitempty"`
}
