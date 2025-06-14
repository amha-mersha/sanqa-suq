package repositories

import (
	"context"
	"errors"

	"github.com/amha-mersha/sanqa-suq/internal"
	"github.com/amha-mersha/sanqa-suq/internal/database"
	"github.com/amha-mersha/sanqa-suq/internal/models"
	"github.com/jackc/pgx/v4"
)

type UserRepository struct {
	DB *database.DB
}

func NewUserRepository(db *database.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) FindUserByID(ctx context.Context, userId string) (*models.User, error) {
	query := `SELECT * FROM users WHERE users.user_id = $1`
	var user models.User

	row := r.DB.Pool.QueryRow(ctx, query, userId)
	if err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.PasswordHash, &user.Phone, &user.Role,
		&user.Provider,
		&user.ProviderID,
		&user.CreatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, internal.NotFound("User not found", err)
		}
		return nil, internal.InternalError("Failed to scan user", err)
	}
	return &user, nil
}

func (r *UserRepository) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT * FROM users WHERE users.email = $1`
	var user models.User
	row := r.DB.Pool.QueryRow(ctx, query, email)
	if err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.PasswordHash, &user.Phone, &user.Role,
		&user.Provider,
		&user.ProviderID,
		&user.CreatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, internal.NotFound("USER_NOT_FOUND", err)
		}
		return nil, internal.InternalError("FAILED_TO_SCAN_USER", err)
	}
	return &user, nil
}

func (r *UserRepository) InsertUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (first_name, last_name, email, password_hash, phone, role, provider, provider_id) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING user_id;`
	err := r.DB.Pool.QueryRow(ctx, query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.PasswordHash,
		user.Phone,
		user.Role,
		user.Provider,
		user.ProviderID).Scan(&user.ID)
	if err != nil {
		return internal.InternalError("Failed to insert user", err)
	}
	return nil
}
