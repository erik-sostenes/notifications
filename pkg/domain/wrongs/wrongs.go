package wrongs

// CommandNotRegisteredError will return an error when an command not registered
type CommandNotRegisteredError string

// CommandNotRegisteredError implements the Error interface
func (e CommandNotRegisteredError) Error() string {
	return string(e)
}

// CommandAlreadyRegisteredError will return an error when an command not registered
type CommandAlreadyRegisteredError string

// CommandAlreadyRegisteredError implements the Error interface
func (e CommandAlreadyRegisteredError) Error() string {
	return string(e)
}

// QueryNotRegisteredError will return an error when an command not registered
type QueryNotRegisteredError string

// QueryNotRegisteredError implements the Error interface
func (e QueryNotRegisteredError) Error() string {
	return string(e)
}

// QueryAlreadyRegisteredError will return an error when an command not registered
type QueryAlreadyRegisteredError string

// QueryAlreadyRegisteredError implements the Error interface
func (e QueryAlreadyRegisteredError) Error() string {
	return string(e)
}

// StatusUnprocessableEntity will return an error when the server cannot process the contained instructions
type StatusUnprocessableEntity string

// StatusBadRequest implements the Error interface
func (e StatusUnprocessableEntity) Error() string {
	return string(e)
}

// StatusBadRequest will return an error when the client makes a mistakes
type StatusBadRequest string

// StatusBadRequest implements the Error interface
func (e StatusBadRequest) Error() string {
	return string(e)
}

// StatusNotFound will return an error when a resource is not found
type StatusNotFound string

// StatusNotFound implements the Error interface
func (e StatusNotFound) Error() string {
	return string(e)
}
