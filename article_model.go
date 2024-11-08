package main

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type ArticleModel struct {
	bun.BaseModel `bun:"table:articles,alias:a"`

	ID        int       `bun:"id,pk,autoincrement"`
	Title     string    `bun:"title,notnull"`
	Contents  string    `bun:"contents,notnull"`
	NiceNum   int       `bun:"nice_num,notnull"`
	CreatedAt time.Time `bun:"created_at,notnull"`
}

func CreateArticleTable(ctx context.Context, db *bun.DB) error {
	_, err := db.NewCreateTable().Model((*ArticleModel)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func TruncateArticleTable(ctx context.Context, db *bun.DB) error {
	if _, err := db.NewTruncateTable().Model((*ArticleModel)(nil)).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func DropArticleTable(ctx context.Context, db *bun.DB) error {
	if _, err := db.NewDropTable().Model((*ArticleModel)(nil)).IfExists().Exec(ctx); err != nil {
		return err
	}

	return nil
}
