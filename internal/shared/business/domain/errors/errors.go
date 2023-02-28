package errors

// StatusBadRequest will return an error when the client makes a mistakes
type StatusBadRequest string

// StatusBadRequest implements the Error interface
func (e StatusBadRequest) Error() string {
	return string(e)
}

// StatusUnprocessableEntity will return an error when the server cannot process the contained instructions
type StatusUnprocessableEntity string

// StatusBadRequest implements the Error interface
func (e StatusUnprocessableEntity) Error() string {
	return string(e)
}

// DuplicatedDomainEvent will return an error when a duplicated domain event is registered
type DuplicatedDomainEvent string

// DuplicatedDomainEvent implements the Error interface
func (e DuplicatedDomainEvent) Error() string {
	return string(e)
}

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
