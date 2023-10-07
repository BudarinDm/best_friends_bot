package logic

import (
	"best_friends_bot/internal/model"
	"context"
)

func (l Logic) GetAdmin(ctx context.Context, id int64) (model.Admin, error) {
	return l.repo.GetAdmin(ctx, id)
}
