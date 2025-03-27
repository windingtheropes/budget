package auth

// Email + ID
type UserIdentifier interface {
	int | string
}