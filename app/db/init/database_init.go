package database_init

import (
	"context"

	"github.com/uptrace/bun"
)

type User struct {
	ID        string `bun:"type:varchar(128),primary" json:"id"`        // ユーザID
	AuthToken string `bun:"type:varchar(128),unique" json:"auth_token"` // 認証トークン
	Name      string `bun:"type:varchar(64)" json:"name"`               // ユーザ名
	Count     int    `bun:"type:int unsigned" json:"count"`             // カウント
}

func CreateTable(db *bun.DB) error {
	_, err := db.NewCreateTable().Model((*User)(nil)).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
