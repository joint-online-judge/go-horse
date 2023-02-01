package schemas

// Pointer returns a pointer to a variable holding the supplied T constant
func Pointer[T any](x T) *T {
	return &x
}

type Pagination struct {
	Ordering *string `form:"ordering,omitempty" json:"ordering,omitempty"`
	Offset   *int    `form:"offset,omitempty"   json:"offset,omitempty"`
	Limit    *int    `form:"limit,omitempty"    json:"limit,omitempty"`
}
