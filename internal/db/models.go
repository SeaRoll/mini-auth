// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"
)

type Account struct {
	ID           string         `json:"id"`
	Email        string         `json:"email"`
	RefreshToken sql.NullString `json:"refreshToken"`
	Enabled      bool           `json:"enabled"`
}
