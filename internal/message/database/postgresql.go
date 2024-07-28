package database

import (
	"context"
	"log"

	"testingwork-kafka/pkg/clients/postresql"
	"time"
)

type repository struct {
	client postresql.Client
}

func NewRepository(client postresql.Client) Repository {
	return &repository{
		client: client,
	}
}

func (r *repository) Message(ctx context.Context, content string) error {
	status := "received"
	q := `INSERT INTO message(content, status, created) VALUES($1, $2, $3)`
	_, err := r.client.Exec(ctx, q, content, status, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		//добавить более тонкий обработчик событий
		log.Fatal("ошибка записи в бд постгрес: ", err)
		return err
	}
	return nil
}

func (r *repository) Statistics(ctx context.Context) int {
	var count int
	q := `SELECT COUNT(*) FROM message WHERE status = 'processed';`
	err := r.client.QueryRow(ctx, q).Scan(&count)
	if err != nil {
		//добавить более тонкий обработчик событий
		log.Fatal("ошибка чтения статистики постгрес: ", err)
	}
	return count
}
