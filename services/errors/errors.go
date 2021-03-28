package errors

type PasswordsDontMatchError struct {}

func (e *PasswordsDontMatchError) Error() string {
	return "Passwords dont match"
}
