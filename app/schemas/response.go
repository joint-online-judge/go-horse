package schemas

type StandardResp[T any] struct {
	BizError
	Data T `json:"data,omitempty"`
}

type EmptyResp StandardResp[*any]

type ListResp[T any] struct {
	Count   int64 `json:"count"`
	Results []T   `json:"results" validate:"dive"`
}

type StandardListResp[T any] StandardResp[ListResp[T]]

type NonStandardResp struct {
	Data any
}
