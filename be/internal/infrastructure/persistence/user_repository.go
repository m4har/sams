package persistence

import (
	"database/sql"
	"time"

	"github.com/m4har/sams/internal/domain/entity"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *entity.User) error {
	query := `
        INSERT INTO users (email, password, created_at, updated_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id`

	now := time.Now()
	return r.db.QueryRow(
		query,
		user.Email,
		user.Password,
		now,
		now,
	).Scan(&user.ID)
}

func (r *UserRepository) FindByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	query := `
        SELECT id, email, password, created_at, updated_at
        FROM users
        WHERE email = $1`

	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
