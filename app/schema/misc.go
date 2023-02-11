package schema

// Version defines model for Version.
type Version struct {
	Git     string `json:"git"`
	Version string `json:"version"`
}
