package seeders

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

var defaultGroupNames = []string{
	"Administrador",
	"Gerente",
	"Usuario",
}

func Groups() Seeder {
	return Seeder{
		Name: "groups",
		Run:  seedGroups,
	}
}

func seedGroups(ctx context.Context, pool *pgxpool.Pool) error {
	const existsQuery = `
		SELECT EXISTS(
			SELECT 1 FROM groups
			WHERE name = $1 AND deleted_at IS NULL
		)
	`
	const insertQuery = `INSERT INTO groups (name) VALUES ($1)`

	for _, raw := range defaultGroupNames {
		name := strings.TrimSpace(raw)
		if name == "" {
			continue
		}

		var exists bool
		if err := pool.QueryRow(ctx, existsQuery, name).Scan(&exists); err != nil {
			return fmt.Errorf("check group %q: %w", name, err)
		}
		if exists {
			continue
		}

		if _, err := pool.Exec(ctx, insertQuery, name); err != nil {
			return fmt.Errorf("insert group %q: %w", name, err)
		}
	}

	return nil
}
