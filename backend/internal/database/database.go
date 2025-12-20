package database

import (
	"context"
	"database/sql"
	"time"
)

type Database struct {
	db     *sql.DB
	logger Logger
}

func New(db *sql.DB, logger Logger) *Database {
	return &Database{
		db:     db,
		logger: logger,
	}
}

func (d *Database) RunQuery(
	ctx context.Context,
	query string,
	mapRow func(*sql.Rows) error,
	args ...any,
) error {

	start := time.Now()

	rows, err := d.db.QueryContext(ctx, query, args...)
	if err != nil {
		d.logger.Error(ctx, "query failed", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		if err := mapRow(rows); err != nil {
			d.logger.Error(ctx, "scan failed", err)
			return err
		}
	}

	d.logger.Info(ctx, "query success",
		"duration", time.Since(start),
	)

	return nil
}

func (d *Database) RunExec(
	ctx context.Context,
	query string,
	args ...any,
) (sql.Result, error) {

	start := time.Now()

	res, err := d.db.ExecContext(ctx, query, args...)
	if err != nil {
		d.logger.Error(ctx, "exec failed", err)
		return nil, err
	}

	d.logger.Info(ctx, "exec success",
		"duration", time.Since(start),
	)

	return res, nil
}
