package errors

type DuplicateEmailError struct {}

func (e *DuplicateEmailError) Error() string {
	return "Duplicate Email"
}