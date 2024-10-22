package models

import (
	"github.com/google/uuid"
	"time"
)

/**
* Shared entities that are imported by more than one package.
**/
type User struct {
	BaseDBDateModel
	Email    string `db:"id" json:"email"`
	Password string `db:"password" json:"password"`
}

/**
* Base models for default table columns.
**/
type BaseDBDateModel struct {
	BaseDBModel
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type BaseDBModel struct {
	ID uuid.UUID `db:"id" json:"id"`
}
