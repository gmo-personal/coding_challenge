package database

import (
	"database/sql"
	"github.com/gmo-personal/coding_challenge/model"
)

// Inserts an account.
func InsertAccount(db *sql.DB, accountInfo *model.Account) error {
	insertAccountStmt := `INSERT INTO account (
		id,
		username,
		password,
		plan
	) VALUES (?, ?, ?, ?);`

	_, err := db.Exec(
		insertAccountStmt,
		accountInfo.ID,
		accountInfo.Username,
		accountInfo.Password,
		accountInfo.Plan)
	return err
}
