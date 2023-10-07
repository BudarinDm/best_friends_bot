package service

import (
	"context"
	"fmt"
)

func (a *App) adminCommandChecker(ctx context.Context, id int64) (bool, error) {
	admin, err := a.logic.GetAdmin(ctx, id)
	if err != nil {
		return false, fmt.Errorf("error GetAdmin: %s", err.Error())
	}
	return admin.Command, nil
}
