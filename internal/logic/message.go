package logic

import (
	"best_friends_bot/internal/model"
	"context"
)

func (l Logic) GetTriggers(ctx context.Context) ([]model.Trigger, error) {
	return l.repo.GetTriggers(ctx)
}
