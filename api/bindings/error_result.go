package bindings

// ErrorResult is used to respond to requests that are not correct
type ErrorResult struct {
	ErrorMessage string `json:"error"`
}
