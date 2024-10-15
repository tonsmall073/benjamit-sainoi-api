package utils

type GenericResult[T any] struct {
	ResModel T
}
type GenericError struct {
	Error      error
	StatusCode int
}
