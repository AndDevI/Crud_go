package seeders

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Seeder struct {
	Name string
	Run  func(ctx context.Context, pool *pgxpool.Pool) error
}

// All registra todos os seeders do projeto.
func All() []Seeder {
	return []Seeder{
		Groups(),
	}
}
