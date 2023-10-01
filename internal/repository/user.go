package db

import (
	"best_friends_bot/internal/model"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type User interface {
	GetBDDates(ctx context.Context) ([]model.UserBirthday, error)
	GetBDDateNext(ctx context.Context, t time.Time) (model.UserBirthday, error)
}

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUser(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) GetBDDates(ctx context.Context) ([]model.UserBirthday, error) {
	query := `
select fio, date_part('month' ,birthday), date_part('day' ,birthday) from public.user
order by date_part('month' ,birthday), date_part('day' ,birthday)
`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error sql query: %w", err)
	}

	var data []model.UserBirthday
	for rows.Next() {
		var t model.UserBirthday
		err = rows.Scan(&t.FIO, &t.Month, &t.Day)
		if err != nil {
			return nil, fmt.Errorf("error sql scan: %w", err)
		}

		t.MonthText = model.BDText[t.Month]
		data = append(data, t)
	}

	return data, nil
}

func (r *UserRepo) GetBDDateNext(ctx context.Context, t time.Time) (model.UserBirthday, error) {
	query := `
select fio, date_part('month' ,birthday), date_part('day' ,birthday) from public.user
where date_part('month' ,birthday) = $1 and date_part('day' ,birthday) >= $2
order by date_part('month' ,birthday), date_part('day' ,birthday)
limit 1
`
	month := int64(t.Month())

	var data model.UserBirthday
	for {
		err := r.db.QueryRow(ctx, query, month, t.Day()).Scan(&data.FIO, &data.Month, &data.Day)
		if err != nil {
			if err == pgx.ErrNoRows {
				if month == 12 {
					month = 1
				} else {
					month += 1
				}
				continue
			}
			return model.UserBirthday{}, fmt.Errorf("error sql query: %w", err)
		}
		break
	}

	return data, nil
}
