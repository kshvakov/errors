package errors

type Error interface {
	Error() string
	Stack() string
}
