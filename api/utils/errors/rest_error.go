package errors

type restErr struct {
	Message []string `json:"message"`
	Status  int      `json:"status"`
	Error   []string
}
