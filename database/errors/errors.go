package errors

type DuplicateEmailError struct {}
func (e *DuplicateEmailError) Error() string {
	return "Duplicate Email"
}

type NoUserFoundError struct {}
func (e *NoUserFoundError) Error() string {
	return "Wrong email"
}

type RecordNotFoundError struct {}
func (e *RecordNotFoundError) Error() string {
	return "Record not found"
}

type DuplicateCategoryError struct {}
func (e *DuplicateCategoryError) Error() string {
	return "Category already exist"
}