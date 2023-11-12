package models

type ComingPrimaryKey struct {
	Id string `json:"id"`
}

type CreateComing struct {
	Remaining_Id string `json:"remaining_id"`
}

type Coming struct {
	Id           string `json:"id"`
	Remaining_Id string `json:"remaining_id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type UpdateComing struct {
	Id           string `json:"id"`
	Remaining_Id string `json:"remaining_id"`
}

type GetListComingRequest struct {
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
	Search string `json:"search"`
}

type GetListComingResponse struct {
	Count   string    `json:"count"`
	Comings []*Coming `json:"comings"`
}
