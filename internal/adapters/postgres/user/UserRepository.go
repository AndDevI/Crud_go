package user

import (
	"context"
	"errors"
	"fmt"

	domain "crmata-go/internal/domain/user"
	userrepository "crmata-go/internal/domain/user/repository"
	"crmata-go/internal/helpers"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) UserRepository {
	return UserRepository{pool: pool}
}

func (r UserRepository) Create(ctx context.Context, u domain.User) (domain.User, error) {
	const query = `
		INSERT INTO users (name, email, telephone, password, image, active, group_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, name, email, password, telephone, image, active, group_id, created_at, updated_at, deleted_at
	`

	row := r.pool.QueryRow(ctx, query,
		u.Name,
		u.Email,
		u.Telephone,
		u.Password,
		u.Image,
		u.Active,
		u.GroupID,
	)

	created, err := scanUserRow(row)
	if err != nil {
		return domain.User{}, mapPostgresError(err)
	}

	return created, nil
}

func (r UserRepository) FindAll(ctx context.Context, filter userrepository.ListFilter) ([]domain.User, error) {
	sortBy := helpers.SafeOrderBy(filter.SortBy, map[string]string{
		"name":       "name",
		"email":      "email",
		"created_at": "created_at",
		"updated_at": "updated_at",
	}, "created_at")
	sortDirection := helpers.SafeSortDirection(filter.SortDirection, "DESC")

	query := `
		SELECT id, name, email, password, telephone, image, active, group_id, created_at, updated_at, deleted_at
		FROM users
		WHERE deleted_at IS NULL
	`
	args := make([]any, 0)

	if filter.Search != "" {
		safeTerm := helpers.EscapeLikeTerm(filter.Search)
		args = append(args, safeTerm)
		idx := len(args)
		query += fmt.Sprintf(
			` AND (name ILIKE '%%' || $%d || '%%' ESCAPE '\\' OR email ILIKE '%%' || $%d || '%%' ESCAPE '\\')`,
			idx,
			idx,
		)
	}

	query += fmt.Sprintf(" ORDER BY %s %s", sortBy, sortDirection)

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]domain.User, 0)
	for rows.Next() {
		u, err := scanUserRow(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r UserRepository) FindByID(ctx context.Context, id int64) (domain.User, error) {
	const query = `
		SELECT id, name, email, password, telephone, image, active, group_id, created_at, updated_at, deleted_at
		FROM users
		WHERE id = $1 AND deleted_at IS NULL
	`

	u, err := scanUserRow(r.pool.QueryRow(ctx, query, id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, userrepository.ErrNotFound
		}
		return domain.User{}, err
	}
	return u, nil
}

func (r UserRepository) Update(ctx context.Context, u domain.User) (domain.User, error) {
	const query = `
		UPDATE users
		SET
			name = $2,
			email = $3,
			telephone = $4,
			password = $5,
			image = $6,
			active = $7,
			group_id = $8,
			updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
		RETURNING id, name, email, password, telephone, image, active, group_id, created_at, updated_at, deleted_at
	`

	updated, err := scanUserRow(r.pool.QueryRow(ctx, query,
		u.ID,
		u.Name,
		u.Email,
		u.Telephone,
		u.Password,
		u.Image,
		u.Active,
		u.GroupID,
	))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, userrepository.ErrNotFound
		}
		return domain.User{}, mapPostgresError(err)
	}

	return updated, nil
}

func (r UserRepository) UpdateGroupID(ctx context.Context, id int64, groupID *int64) (domain.User, error) {
	const query = `
		UPDATE users
		SET
			group_id = $2,
			updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
		RETURNING id, name, email, password, telephone, image, active, group_id, created_at, updated_at, deleted_at
	`

	updated, err := scanUserRow(r.pool.QueryRow(ctx, query, id, groupID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, userrepository.ErrNotFound
		}
		return domain.User{}, mapPostgresError(err)
	}

	return updated, nil
}

func (r UserRepository) Delete(ctx context.Context, id int64) error {
	const query = `
		UPDATE users
		SET deleted_at = NOW(), updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`

	result, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return userrepository.ErrNotFound
	}
	return nil
}

func scanUserRow(row pgx.Row) (domain.User, error) {
	var u domain.User
	var telephone *string
	var image *string
	var groupID *int64

	err := row.Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.Password,
		&telephone,
		&image,
		&u.Active,
		&groupID,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.DeletedAt,
	)
	if err != nil {
		return domain.User{}, err
	}

	u.Telephone = telephone
	u.Image = image
	u.GroupID = groupID
	return u, nil
}

func mapPostgresError(err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return userrepository.ErrNotFound
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return userrepository.ErrEmailAlreadyExists
		case "23503":
			return userrepository.ErrGroupNotFound
		}
	}

	return err
}
