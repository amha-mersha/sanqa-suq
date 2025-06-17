package repositories

import (
	"context"
	"errors"

	"github.com/amha-mersha/sanqa-suq/internal/database"
	errs "github.com/amha-mersha/sanqa-suq/internal/errors"
	"github.com/amha-mersha/sanqa-suq/internal/models"
	"github.com/jackc/pgx/v5"
)

type AddressRepository struct {
	DB *database.DB
}

func NewAddressRepository(db *database.DB) *AddressRepository {
	return &AddressRepository{DB: db}
}

func (r *AddressRepository) CreateAddress(ctx context.Context, address *models.Address) error {
	query := `
		INSERT INTO addresses (user_id, street, city, state, postal_code, country, address_type)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING address_id`

	err := r.DB.Pool.QueryRow(
		ctx,
		query,
		address.UserID,
		address.Street,
		address.City,
		address.State,
		address.PostalCode,
		address.Country,
		address.Type,
	).Scan(&address.AddressID)

	if err != nil {
		return errs.InternalError("failed to create address", err)
	}

	return nil
}

func (r *AddressRepository) GetAddressByID(ctx context.Context, addressID int, userID string) (*models.Address, error) {
	query := `
		SELECT address_id, user_id, street, city, state, postal_code, country, address_type
		FROM addresses
		WHERE address_id = $1 AND user_id = $2`

	address := &models.Address{}
	err := r.DB.Pool.QueryRow(ctx, query, addressID, userID).Scan(
		&address.AddressID,
		&address.UserID,
		&address.Street,
		&address.City,
		&address.State,
		&address.PostalCode,
		&address.Country,
		&address.Type,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errs.NotFound("address not found", err)
		}
		return nil, errs.InternalError("failed to get address", err)
	}

	return address, nil
}

func (r *AddressRepository) GetUserAddresses(ctx context.Context, userID string) ([]*models.Address, error) {
	query := `
		SELECT address_id, user_id, street, city, state, postal_code, country, address_type
		FROM addresses
		WHERE user_id = $1
		ORDER BY address_id DESC`

	rows, err := r.DB.Pool.Query(ctx, query, userID)
	if err != nil {
		return nil, errs.InternalError("failed to get user addresses", err)
	}
	defer rows.Close()

	var addresses []*models.Address
	for rows.Next() {
		address := &models.Address{}
		err := rows.Scan(
			&address.AddressID,
			&address.UserID,
			&address.Street,
			&address.City,
			&address.State,
			&address.PostalCode,
			&address.Country,
			&address.Type,
		)
		if err != nil {
			return nil, errs.InternalError("failed to scan address", err)
		}
		addresses = append(addresses, address)
	}

	if err := rows.Err(); err != nil {
		return nil, errs.InternalError("error iterating addresses", err)
	}

	return addresses, nil
}

func (r *AddressRepository) UpdateAddress(ctx context.Context, address *models.Address) error {
	query := `
		UPDATE addresses
		SET street = $1, city = $2, state = $3, postal_code = $4, country = $5, address_type = $6
		WHERE address_id = $7 AND user_id = $8`

	result, err := r.DB.Pool.Exec(
		ctx,
		query,
		address.Street,
		address.City,
		address.State,
		address.PostalCode,
		address.Country,
		address.Type,
		address.AddressID,
		address.UserID,
	)
	if err != nil {
		return errs.InternalError("failed to update address", err)
	}

	if result.RowsAffected() == 0 {
		return errs.NotFound("address not found", nil)
	}

	return nil
}

func (r *AddressRepository) DeleteAddress(ctx context.Context, addressID int, userID string) error {
	query := `DELETE FROM addresses WHERE address_id = $1 AND user_id = $2`

	result, err := r.DB.Pool.Exec(ctx, query, addressID, userID)
	if err != nil {
		return errs.InternalError("failed to delete address", err)
	}

	if result.RowsAffected() == 0 {
		return errs.NotFound("address not found", nil)
	}

	return nil
}
