package http

// ErrorMessage is a json serializable error.
type ErrorMessage struct {
	Message string `json:"message"`
}

func (e *ErrorMessage) Error() string {
	return e.Message
}

// NewErrorMessage creates a new NewErrorMessage.
func NewErrorMessage(e error) *ErrorMessage {
	return &ErrorMessage{e.Error()}
}
