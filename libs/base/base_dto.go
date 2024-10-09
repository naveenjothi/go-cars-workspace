package base

type BaseDto[T any] struct {
	Items []T   `json:"items"`
	Count int32 `json:"count"`
}
