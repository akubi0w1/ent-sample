package msyql

import (
	"context"
	"testing"
	"time"

	"github.com/akubi0w1/ent-sample/ent/enttest"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestUser_Create(t *testing.T) {
	ctx := context.Background()
	t.Run("success", func(t *testing.T) {
		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
		defer client.Close()

		dummyTime := time.Date(2021, 4, 15, 0, 0, 0, 0, time.UTC)
		expected := 1

		out, err := Create(client, ctx, "accID", dummyTime)
		assert.Equal(t, expected, out)
		assert.Nil(t, err)
	})
}
