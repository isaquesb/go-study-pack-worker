package domain

const (
	TargetPublic string = "public"
	TargetPrivate string = "private"
)

type Target struct {
	Driver     TargetDriver `json:"driver"`
	Visibility string       `json:"visibility"`
}

type TargetDriver struct {
	Requires []flag `json:"requirements"`
	Desired []flag  `json:"desired"`
}

type flag struct {
	Property string `json:"property"`
	Value string `json:"value"`
}
