package database

import "context"

type Repository interface {
	Message(ctx context.Context, content, status string) error
	Statistics(ctx context.Context, conent Contents) (string, error)
}
