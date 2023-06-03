package input

type UserHandlerInput struct {
	PageNum  int64 `json:"page_num"`
	PageSize int64 `json:"page_size"`
}
