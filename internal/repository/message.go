package db

import (
	"best_friends_bot/internal/model"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Message interface {
	GetTriggers(ctx context.Context) ([]model.Trigger, error)
	GetMessageByTrigger(ctx context.Context, trigger string) (message []model.Message, err error)
}

type MessageRepo struct {
	db *pgxpool.Pool
}

func NewMessage(db *pgxpool.Pool) *MessageRepo {
	return &MessageRepo{db: db}
}

func (r *MessageRepo) GetTriggers(ctx context.Context) ([]model.Trigger, error) {
	query := `select id, trigger from trigger`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error sql query: %w", err)
	}

	var triggers []model.Trigger
	for rows.Next() {
		var t model.Trigger
		err = rows.Scan(&t.Id, &t.Trigger)
		if err != nil {
			return nil, fmt.Errorf("error sql scan: %w", err)
		}
		triggers = append(triggers, t)
	}

	return triggers, nil
}

func (r *MessageRepo) GetMessageByTrigger(ctx context.Context, trigger string) (message []model.Message, err error) {
	query := `
SELECT t.id, m.message FROM trigger_message tm
JOIN trigger t ON tm.trigger_id = t.id
JOIN message m ON tm.message_id = m.id
WHERE t.trigger = $1
`

	rows, err := r.db.Query(ctx, query, trigger)
	if err != nil {
		return nil, fmt.Errorf("error sql query: %w", err)
	}
	for rows.Next() {
		var t model.Message
		err = rows.Scan(&t.Id, &t.Message)
		if err != nil {
			return nil, fmt.Errorf("error sql scan: %w", err)
		}
		message = append(message, t)
	}
	return
}
