package types

type StandardResp[T any] struct {
	BizError
	Data T `json:"data,omitempty"`
}

type EmptyResp StandardResp[*any]
