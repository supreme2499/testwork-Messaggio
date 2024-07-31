package database

import "context"

type Repository interface {
	Message(ctx context.Context, content string) error
	MessageWork(ctx context.Context, content string) error
	Statistics(ctx context.Context) int
}
