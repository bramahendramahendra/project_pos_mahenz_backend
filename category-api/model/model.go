package model

type Category struct {
	ID        uint64  `json:"id"`
	Category  *string `json:"category"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at"`
	DeletedAt *string `json:"deleted_at,omitempty"`
}

type Response struct {
	ResponseCode int           `json:"responseCode"`
	ResponseDesc string        `json:"responseDesc"`
	ResponseData interface{}   `json:"responseData"`
	ResponseMeta *ResponseMeta `json:"responseMeta,omitempty"`
}

type ResponseMeta struct {
	Page         int `json:"page"`
	PerPage      int `json:"per_page"`
	TotalPages   int `json:"total_pages"`
	TotalRecords int `json:"total_records"`
}

type CategoryInput struct {
	Category string `json:"category" binding:"required" example:"Makanan"`
}
