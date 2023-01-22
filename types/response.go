package types

type StandardResp[T any] struct {
	BizError
	Data T `json:"data,omitempty"`
}

type EmptyResp StandardResp[*any]

type ListResp[T any] struct {
	Count   int `json:"count"`
	Results []T `json:"results" validate:"dive"`
}

type StandardListResp[T any] StandardResp[ListResp[T]]
