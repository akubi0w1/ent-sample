package msyql

import (
	"context"
	"time"

	"github.com/akubi0w1/ent-sample/ent"
)

func Create(cli *ent.Client, ctx context.Context, accountID string, now time.Time) (int, error) {
	u, err := cli.User.Create().
		SetAccountID("accountID").
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)
	if err != nil {
		return 0, err
	}
	return u.ID, nil
}
