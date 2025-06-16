package repositories

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/amha-mersha/sanqa-suq/internal/database"
	errs "github.com/amha-mersha/sanqa-suq/internal/errors"
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
			return nil, errs.NotFound("User not found", err)
		}
		return nil, errs.InternalError("Failed to scan user", err)
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
			return nil, errs.NotFound("USER_NOT_FOUND", err)
		}
		return nil, errs.InternalError("FAILED_TO_SCAN_USER", err)
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
		return errs.InternalError("Failed to insert user", err)
	}
	return nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, userId string, updateFields map[string]any) (*models.User, error) {
	setClauses := []string{}
	values := []any{}
	i := 1
	for field, value := range updateFields {
		setClauses = append(setClauses, field+" = $"+strconv.Itoa(i))
		values = append(values, value)
		i++
	}
	values = append(values, userId)
	query := fmt.Sprintf(`
		UPDATE users 
		SET %s 
		WHERE user_id = $%d 
		RETURNING user_id, email, first_name, last_name, phone, role, created_at, provider, provider_id
	`, strings.Join(setClauses, ", "), i)

	var updatedUser models.User
	err := r.DB.Pool.QueryRow(ctx, query, values...).Scan(
		&updatedUser.ID,
		&updatedUser.Email,
		&updatedUser.FirstName,
		&updatedUser.LastName,
		&updatedUser.Phone,
		&updatedUser.Role,
		&updatedUser.CreatedAt,
		&updatedUser.Provider,
		&updatedUser.ProviderID,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errs.NotFound("User not found", err)
		}
		return nil, errs.InternalError("Failed to update user", err)
	}
	return &updatedUser, nil
}
