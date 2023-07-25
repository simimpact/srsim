package gen

//nolint:tagliatelle // need to match datamine
type HashInfo struct {
	Hash int `json:"Hash"`
}

//nolint:tagliatelle // need to match datamine
type ValueInfo struct {
	Value float64 `json:"Value"`
}

//nolint:tagliatelle // need to match datamine
type StatusAdd struct {
	PropertyType string    `json:"PropertyType"`
	Value        ValueInfo `json:"Value"`
}

//nolint:tagliatelle // need to match datamine
type TargetInfo struct {
	TargetType string `json:"TargetType"`
}
