package model

type Product struct {
	ID         uint64  `json:"id"`
	IDCategory uint64  `json:"id_category"`
	Product    *string `json:"product"`
	CreatedAt  *string `json:"created_at,omitempty"`
	UpdatedAt  *string `json:"updated_at"`
	DeletedAt  *string `json:"deleted_at,omitempty"`
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

type ProductInput struct {
	IDCategory uint64 `json:"id_category" binding:"required" example:"1"`
	Product    string `json:"product" binding:"required" example:"Chitato"`
}
