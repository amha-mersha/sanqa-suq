package repositories

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/amha-mersha/sanqa-suq/internal/database"
	"github.com/amha-mersha/sanqa-suq/internal/dtos"
	errs "github.com/amha-mersha/sanqa-suq/internal/errors"
	"github.com/amha-mersha/sanqa-suq/internal/models"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v5"
)

type ReviewRepository struct {
	DB *database.DB
}

func NewReviewRepository(db *database.DB) *ReviewRepository {
	return &ReviewRepository{DB: db}
}

func (repository *ReviewRepository) CreateNewReview(ctx context.Context, reviewDTO *dtos.CreateReviewDTO) (*models.Review, error) {
	query := `
		INSERT INTO reviews (
			user_id,
			product_id,
			rating,
			comment
		) VALUES (
			$1, $2, $3, $4
		)
		RETURNING review_id, user_id, product_id, rating, comment, review_date;
	`

	newReview := &models.Review{}

	err := repository.DB.Pool.QueryRow(
		ctx,
		query,
		reviewDTO.UserID,
		reviewDTO.ProductID,
		reviewDTO.Rating,
		reviewDTO.Comment,
	).Scan(
		&newReview.ReviewId,
		&newReview.UserId,
		&newReview.ProductId,
		&newReview.Rating,
		&newReview.Comment,
		&newReview.ReviewDate,
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23503" {
				return nil, errs.BadRequest("FOREIGN_KEY_VIOLATION", fmt.Errorf("invalid user_id or product_id: %w", err))
			}
		}
		if err == pgx.ErrNoRows {
			return nil, errs.InternalError("failed to create a new review: no rows returned", err)
		}
		return nil, errs.InternalError("failed to create a review", err)
	}

	return newReview, nil
}

func (repository *ReviewRepository) UpdateReview(
	ctx context.Context,
	reviewID string,
	dto *dtos.UpdateReviewDTO,
) (*models.Review, error) {
	fields := map[string]any{}

	if dto.Rating != nil {
		fields["rating"] = *dto.Rating
	}
	if dto.Comment != nil {
		fields["comment"] = *dto.Comment
	}

	if len(fields) == 0 {
		return nil, errs.BadRequest("NO_FIELDS_TO_UPDATE", errors.New("no fields provided for update"))
	}

	setClauses := []string{}
	args := []any{}
	i := 1

	for col, val := range fields {

		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", col, i))
		args = append(args, val)
		i++
	}

	args = append(args, reviewID)
	query := fmt.Sprintf(`
		UPDATE reviews
		SET %s
		WHERE review_id = $%d
		RETURNING review_id, user_id, product_id, rating, comment, review_date
	`, strings.Join(setClauses, ", "), i)

	review := &models.Review{}
	err := repository.DB.Pool.QueryRow(ctx, query, args...).Scan(
		&review.ReviewId,
		&review.UserId,
		&review.ProductId,
		&review.Rating,
		&review.Comment,
		&review.ReviewDate,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errs.NotFound(fmt.Sprintf("review with id %s not found", reviewID), err)
		}
		return nil, errs.InternalError(fmt.Sprintf("failed to update review with id %s", reviewID), err)
	}

	return review, nil
}

func (repository *ReviewRepository) FindReviewByID(ctx context.Context, reviewID string) (*models.Review, error) {
	var review models.Review

	query := `
		SELECT review_id, user_id, product_id, rating, comment, review_date
		FROM reviews
		WHERE review_id = $1
	`

	err := repository.DB.Pool.QueryRow(ctx, query, reviewID).Scan(
		&review.ReviewId,
		&review.UserId,
		&review.ProductId,
		&review.Rating,
		&review.Comment,
		&review.ReviewDate,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errs.NotFound(fmt.Sprintf("review with id %s not found", reviewID), err)
		}
		return nil, errs.InternalError(fmt.Sprintf("failed to find review with id %s", reviewID), err)
	}

	return &review, nil
}

func (repository *ReviewRepository) DeleteReviewByID(ctx context.Context, reviewID string) error {
	query := `DELETE FROM reviews WHERE review_id = $1`

	result, err := repository.DB.Pool.Exec(ctx, query, reviewID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) || result.RowsAffected() == 0 {
			return errs.NotFound(fmt.Sprintf("review with id %s not found", reviewID), err)
		}
		return errs.InternalError(fmt.Sprintf("failed to delete review with id %s", reviewID), err)
	}

	return nil
}
