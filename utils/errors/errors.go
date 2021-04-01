package errors

type MaximumCategoriesError struct {}
func (e *MaximumCategoriesError) Error() string {
	return "Maximum categories"
}
