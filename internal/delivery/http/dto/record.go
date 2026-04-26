package dto

type CreateRecordRequest struct {
	Name string `json:"name"`
}

type UpdateRecordRequest struct {
	Name string `json:"name"`
}

type RecordResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
