package database

import (
	"context"
	"fmt"

	"crmata-go/internal/infrastructure/database/seeders"
	"github.com/jackc/pgx/v5/pgxpool"
)

// RunSeeders executa todos os seeders registrados.
func RunSeeders(ctx context.Context, pool *pgxpool.Pool) error {
	for _, seeder := range seeders.All() {
		if err := seeder.Run(ctx, pool); err != nil {
			return fmt.Errorf("seed %s: %w", seeder.Name, err)
		}
	}
	return nil
}
