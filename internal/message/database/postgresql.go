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

func (r *repository) Message(ctx context.Context, content, status string) error {
	q := `INSERT INTO mescont(content, status, created) VALUES($1, $2, $3) RETURNING id`
	_, err := r.client.Exec(ctx, q, content, status, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		//добавить более тонкий обработчик событий
		log.Fatal("ошибка записи в бд: ", err)
		return err
	}
	return nil
}

func (r *repository) Statistics(ctx context.Context, conent Contents) (string, error) {

	return "", nil
}
