package schemas

// Pointer returns a pointer to a variable holding the supplied T constant
func Pointer[T any](x T) *T {
	return &x
}
