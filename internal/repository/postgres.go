package repository

import (
	"context"
	"fmt"

	"github.com/Pyth0nHater/link-shorter/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresLinkRepository struct {
	db *pgxpool.Pool
}

func NewPostgresLinkRepository(db *pgxpool.Pool) LinkRepositoryInterface {
	return &PostgresLinkRepository{db: db}
}

func (r *PostgresLinkRepository) CreateLink(ctx context.Context, createLink *models.CreateLink) error {
	query := `
		INSERT INTO links (main_link, short_link) 
		VALUES ($1, $2)
		ON CONFLICT (short_link) DO UPDATE
        SET main_link = EXCLUDED.main_link
	`
	
	_, err := r.db.Exec(
		ctx,
		query,
		createLink.MainLink,
		createLink.ShortLink,
	)

	if err != nil {
		return err
	}

	return nil
}


func (r *PostgresLinkRepository) GetLink(ctx context.Context, getLink models.GetLinkFilter) (*models.GetLink, error) {
	
	if getLink.ShortLink == nil && getLink.MainLink == nil{
		return nil, ErrInvalidGetFilter
	}

	query := `
		SELECT main_link, short_link FROM links 
		WHERE TRUE
	`

	args := make([]any, 0) 
	counter := 1
		
	if getLink.MainLink != nil{
			query+=fmt.Sprintf(` AND main_link = $%d`, counter)
			counter++
			args = append(args, *getLink.MainLink)
	} 

	if getLink.ShortLink != nil{
		query+=fmt.Sprintf(` AND short_link = $%d`, counter)
		counter++
		args = append(args, *getLink.ShortLink)
	}

	var res models.GetLink
	err := r.db.QueryRow(
		ctx,
		query,
		args...
	).Scan(&res.MainLink, &res.ShortLink)

	
	if err != nil {
		return nil, err
	}

	return &res, nil
}

