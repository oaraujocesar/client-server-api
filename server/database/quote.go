package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/oaraujocesar/client-server-api/server/types"
)

func InsertQuote(ctx context.Context, db *sql.DB, quote *types.Quote) error {
	// Define 10ms timeout
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	stmt, err := db.Prepare("INSERT INTO quotes (bid) VALUES (?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, quote.Data.Bid)
	if err != nil {
		return err
	}

	return nil
}
